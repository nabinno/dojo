package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

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
