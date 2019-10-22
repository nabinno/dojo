// @see
//   - https://github.com/awslabs/aws-apigateway-lambda-authorizer-blueprints/blob/master/blueprints/go/main.go
//   - https://github.com/aws/aws-lambda-go/blob/master/events/README_ApiGatewayCustomAuthorizer.md
package main

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	log.Println("Client token: " + event.AuthorizationToken)
	log.Println("Method ARN: " + event.MethodArn)

	methodArnTmp := strings.Split(event.MethodArn, ":")
	apiGatewayArnTmp := strings.Split(methodArnTmp[5], "/")

	resp := NewAuthorizerResponse(
		"user|a1b2c3d4", // principalID
		methodArnTmp[4], // awsAccountID
	)
	resp.Region = methodArnTmp[3]
	resp.APIID = apiGatewayArnTmp[0]
	resp.Stage = apiGatewayArnTmp[1]
	resp.DenyAllMethods()
	resp.AllowMethod(Get, "/pets/*")

	// @note
	//   Add additional key-value pairs associated with the authenticated principal these are made available by APIGW
	//   like so: $context.authorizer.<key> additional context is cached
	resp.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}

	return resp.APIGatewayCustomAuthorizerResponse, nil
}

func main() {
	lambda.Start(handleRequest)
}

// HTTPVerb ....
type HTTPVerb int

const (
	// Get ....
	Get HTTPVerb = iota
	// Post ....
	Post
	// Put ....
	Put
	// Delete ....
	Delete
	// Patch ....
	Patch
	// Head ....
	Head
	// Options ....
	Options
	// All ....
	All
)

func (hv HTTPVerb) String() string {
	switch hv {
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Delete:
		return "DELETE"
	case Patch:
		return "PATCH"
	case Head:
		return "HEAD"
	case Options:
		return "OPTIONS"
	case All:
		return "*"
	}
	return ""
}

// Effect ....
type Effect int

const (
	// Allow ....
	Allow Effect = iota
	// Deny ....
	Deny
)

func (e Effect) String() string {
	switch e {
	case Allow:
		return "Allow"
	case Deny:
		return "Deny"
	}
	return ""
}

// AuthorizerResponse ....
type AuthorizerResponse struct {
	events.APIGatewayCustomAuthorizerResponse
	// The region where the API is deployed. By default this is set to '*'
	Region string
	// The AWS account id the policy will be generated for. This is used to create the method ARNs.
	AccountID string
	// The API Gateway API id. By default this is set to '*'
	APIID string
	// The name of the stage used in the policy. By default this is set to '*'
	Stage string
}

// NewAuthorizerResponse ....
// @param {string} principalID
//   this could be accomplished in a number of ways:
//   1. Call out to OAuth provider
//   2. Decode a JWT token inline
//   3. Lookup in a self-managed DB
// @param {string} AccountID
func NewAuthorizerResponse(principalID string, AccountID string) *AuthorizerResponse {
	return &AuthorizerResponse{
		APIGatewayCustomAuthorizerResponse: events.APIGatewayCustomAuthorizerResponse{
			PrincipalID: principalID,
			PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
			},
		},
		Region:    "*",
		AccountID: AccountID,
		APIID:     "*",
		Stage:     "*",
	}
}

func (r *AuthorizerResponse) addMethod(effect Effect, verb HTTPVerb, resource string) {
	resourceArn := "arn:aws:execute-api:" +
		r.Region + ":" +
		r.AccountID + ":" +
		r.APIID + "/" +
		r.Stage + "/" +
		verb.String() + "/" +
		strings.TrimLeft(resource, "/")

	s := events.IAMPolicyStatement{
		Effect:   effect.String(),
		Action:   []string{"execute-api:Invoke"},
		Resource: []string{resourceArn},
	}

	r.PolicyDocument.Statement = append(r.PolicyDocument.Statement, s)
}

// AllowAllMethods ....
func (r *AuthorizerResponse) AllowAllMethods() {
	r.addMethod(Allow, All, "*")
}

// DenyAllMethods ....
func (r *AuthorizerResponse) DenyAllMethods() {
	r.addMethod(Deny, All, "*")
}

// AllowMethod ....
func (r *AuthorizerResponse) AllowMethod(verb HTTPVerb, resource string) {
	r.addMethod(Allow, verb, resource)
}

// DenyMethod ....
func (r *AuthorizerResponse) DenyMethod(verb HTTPVerb, resource string) {
	r.addMethod(Deny, verb, resource)
}
