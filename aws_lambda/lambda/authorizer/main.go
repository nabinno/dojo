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
	resp.AllowMethod(All, "/pets/*")

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
