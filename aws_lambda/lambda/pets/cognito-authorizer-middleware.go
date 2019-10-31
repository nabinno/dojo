package main

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ErrorMessage is error message struct
type ErrorMessage struct {
	Text string `json:"text"`
}

// JWK is JSON data struct for JSON Web Key
type JWK struct {
	Keys []JWKKey
}

// JWKKey is JSON data struct for Cognito JWK key
type JWKKey struct {
	Alg string
	E   string
	Kid string
	Kty string
	N   string
	Use string
}

// CognitoAuthorizationMiddleware validates Amazon Cognito User Pool JWT or Google JWT
func CognitoAuthorizationMiddleware(userPoolID string) gin.HandlerFunc {
	// @note 1. Download and store the JSON Web Key (JWK) for your User Pool.
	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json", region, userPoolID)
	log.Println(jwkURL)
	jwkMap := getJWKMap(jwkURL)

	// Add Google JWK not for duplicating kid.
	for key, val := range getJWKMap("https://www.googleapis.com/oauth2/v3/certs") {
		jwkMap[key] = val
	}
	log.Println(jwkMap)

	return func(c *gin.Context) {
		tokenString, ok := getBearer(c.Request.Header[authorizationHeaderName])
		if !ok {
			c.AbortWithStatusJSON(401, ErrorMessage{Text: "Authorization Bearer Header is missing"})
			return
		}

		token, err := ValidateToken(tokenString, userPoolID, jwkMap)
		if err != nil || !token.Valid {
			fmt.Printf("token is not valid\n%v", err)
			c.AbortWithStatusJSON(401, ErrorMessage{Text: fmt.Sprintf("token is not valid%v", err)})
		} else {
			// Setting authenticated token, then resource paths can use the user information
			c.Set("token", token)
			c.Next()
		}
	}
}

func getJWKMap(jwkURL string) map[string]JWKKey {
	jwk := &JWK{}
	err := convertURLToJWK(jwkURL, jwk)
	if err != nil {
		panic(err)
	}

	jwkMap := make(map[string]JWKKey, 0)
	for _, jwk := range jwk.Keys {
		jwkMap[jwk.Kid] = jwk
	}

	return jwkMap
}

func convertURLToJWK(url string, target interface{}) error {
	var client = &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		cerr := r.Body.Close()
		if err == nil {
			err = cerr
		}
	}()
	if err != nil {
		return err
	}

	return json.NewDecoder(r.Body).Decode(target)
}

func getBearer(auth []string) (jwt string, ok bool) {
	for _, v := range auth {
		ret := strings.Split(v, " ")
		if len(ret) > 1 && ret[0] == "Bearer" {
			return ret[1], true
		}
	}
	return "", false
}

// ValidateToken validates Amazon Cognito User Pool JWT or Google JWT
func ValidateToken(tokenStr, userPoolID string, jwk map[string]JWKKey) (*jwt.Token, error) {
	// @note 2. Decode the token string into JWT format.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Methods both of Cognito User Pools and Google are RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// @note 5. Get the kid from the JWT token header and retrieve the corresponding JSON Web Key that was stored
		if kid, ok := token.Header["kid"]; ok {
			if kidStr, ok := kid.(string); ok {
				key := jwk[kidStr]
				// @note 6. Verify the signature of the decoded JWT token.
				rsaPublicKey := convertKey(key.E, key.N)
				return rsaPublicKey, nil
			}
		}
		// Does not get RSA public key
		return "", nil
	})
	if err != nil {
		return token, err
	}

	claims := token.Claims.(jwt.MapClaims)
	iss, ok := claims["iss"]
	if !ok {
		return token, fmt.Errorf("token does not contain issuer")
	}

	issStr := iss.(string)
	if strings.Contains(issStr, "cognito-idp") {
		err = validateAWSJWTClaims(claims, userPoolID)
		if err != nil {
			return token, err
		}
	} else if strings.Contains(issStr, "accounts.google.com") {
		err = validateGoogleJWTClaims(claims)
		if err != nil {
			return token, err
		}
	}

	if token.Valid {
		return token, nil
	}

	return token, err
}

// @see https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-identity-user-pools-using-id-and-access-tokens-in-web-api
func validateAWSJWTClaims(claims jwt.MapClaims, userPoolID string) error {
	var err error
	// @note 3. Check the iss claim. It should match your User Pool.
	issShoudBe := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v", region, userPoolID)
	err = validateClaimItem("iss", []string{issShoudBe}, claims)
	if err != nil {
		return err
	}

	// @note 4. Check the token_use claim.
	validateTokenUse := func() error {
		if tokenUse, ok := claims["token_use"]; ok {
			if tokenUseStr, ok := tokenUse.(string); ok {
				if tokenUseStr == "id" || tokenUseStr == "access" {
					return nil
				}
			}
		}
		return errors.New("token_use should be id or access")
	}

	err = validateTokenUse()
	if err != nil {
		return err
	}

	// @note 7. Check the exp claim and make sure the token is not expired.
	err = validateExpired(claims)
	if err != nil {
		return err
	}

	return nil
}

// @see https://developers.google.com/identity/sign-in/web/backend-auth
func validateGoogleJWTClaims(claims jwt.MapClaims) error {
	var err error
	issShoudBe := []string{"accounts.google.com", "https://accounts.google.com"}
	err = validateClaimItem("iss", issShoudBe, claims)
	if err != nil {
		return err
	}

	aud := []string{os.Getenv("GOOGLE_CLIENT_ID")}
	err = validateClaimItem("aud", aud, claims)
	if err != nil {
		return err
	}

	err = validateExpired(claims)
	if err != nil {
		return err
	}

	return nil
}

func validateClaimItem(key string, keyShouldBe []string, claims jwt.MapClaims) error {
	if val, ok := claims[key]; ok {
		if valStr, ok := val.(string); ok {
			for _, shouldbe := range keyShouldBe {
				if valStr == shouldbe {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("%v does not match any of valid values: %v", key, keyShouldBe)
}

func validateExpired(claims jwt.MapClaims) error {
	if tokenExp, ok := claims["exp"]; ok {
		if exp, ok := tokenExp.(float64); ok {
			now := time.Now().Unix()
			fmt.Printf("current unixtime : %v\n", now)
			fmt.Printf("expire unixtime  : %v\n", int64(exp))
			if int64(exp) > now {
				return nil
			}
		}
		return errors.New("cannot parse token exp")
	}
	return errors.New("token is expired")
}

// @see https://gist.github.com/MathieuMailhos/361f24316d2de29e8d41e808e0071b13
func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}

	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}

	pubKey.N.SetBytes(decodedN)
	// log.Println(decodedN)
	// log.Println(decodedE)
	// log.Printf("%#v\n", *pubKey)

	return pubKey
}
