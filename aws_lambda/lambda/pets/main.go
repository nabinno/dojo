package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
)

var (
	envRegion                  string = os.Getenv("REGION")
	envAllowedOrigin           string = os.Getenv("ALLOWED_ORIGIN")
	envAdminsGroupName         string = os.Getenv("ADMINS_GROUP_NAME") // @todo 2019-10-31
	envUsersGroupName          string = os.Getenv("USERS_GROUP_NAME")  // @todo 2019-10-31
	envUserPoolID              string = os.Getenv("USER_POOL_ID")
	envAuthorizationHeaderName string = os.Getenv("AUTHORIZATION_HEADER_NAME")

	ginLambda  *ginadapter.GinLambda
	dynamodb   = dynamo.New(session.New(&aws.Config{Region: aws.String(envRegion)}))
	itemsTable = dynamodb.Table(os.Getenv("ITEMS_TABLE_NAME"))
	usersTable = dynamodb.Table(os.Getenv("USERS_TABLE_NAME")) // @todo 2019-10-31
	// cognito = aws.CognitoIdentityServiceProvider.New()
	// forceSignOutHandler ForceSignOutHandler // ?
)

// Handler is the main entry point for Lambda. Receives a proxy request and returns a proxy response
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")

		r := gin.Default()
		r.Use(CORSMiddleware())

		auth := r.Group("/", CognitoAuthorizationMiddleware())
		auth.GET("/pets", getPets)
		auth.GET("/pets/:id", getPet)
		auth.POST("/pets", createPet)
		// r.GET("/pets", getPets)
		// r.GET("/pets/:id", getPet)
		// r.POST("/pets", createPet)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

// GET /pets
func getPets(c *gin.Context) {
	limit := 10
	if c.Query("limit") != "" {
		newLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 10
		} else {
			limit = newLimit
		}
	}
	if limit > 50 {
		limit = 50
	}
	var pets = make([]Pet, limit)

	for i := 0; i < limit; i++ {
		pets[i] = getRandomPet()
	}

	c.JSON(200, pets)
}

// GET /pets/:id
func getPet(c *gin.Context) {
	// token := c.MustGet("token")
	// claims := token.(*jwt.Token).Claims.(jwt.MapClaims)
	// user := make([]string, 0)
	// if email, ok := claims["email"]; ok {
	//	log.Println(email)
	//	user = append(user, email.(string))
	// }
	// if username, ok := claims["cognito:username"]; ok {
	//	log.Println(username)
	//	user = append(user, username.(string))
	// }

	var petID = c.Param("id")
	var pet Pet

	err := itemsTable.Get("ID", petID).One(&pet)
	if err != nil {
		panic(err)
	}

	c.JSON(200, pet)
}

// POST /pets
func createPet(c *gin.Context) {
	var newPet = Pet{}

	err := c.BindJSON(&newPet)
	if err != nil {
		return
	}
	// newPet.ID = getUUID()

	err = itemsTable.Put(newPet).Run()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusAccepted, newPet)
}
