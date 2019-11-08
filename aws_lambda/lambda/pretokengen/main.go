package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {
	log.Printf("PreTokenGen of user: %s\n", event.UserName)

	// @todo 2019-11-03
	event.Response.ClaimsOverrideDetails.ClaimsToSuppress = []string{"family_name"}

	return event, nil
}

func main() {
	lambda.Start(handleRequest)
}
