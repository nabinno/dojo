package main

import "strings"

// MethodArn ....
type MethodArn struct {
	Region        string
	AwsAccount    string
	APIGatewayArn string
}

// NewMethodArn ....
func NewMethodArn(methodArn string) (rc MethodArn) {
	ss := strings.Split(methodArn, ":")
	rc.Region = ss[3]
	rc.AwsAccount = ss[4]
	rc.APIGatewayArn = ss[5]
	return rc
}

// APIGatewayArn ....
type APIGatewayArn struct {
	APIID    string
	Stage    string
	Method   string
	Resource string
}

// NewAPIGatewayArn ....
func NewAPIGatewayArn(apiGatewayArn string) (rc APIGatewayArn) {
	ss := strings.Split(apiGatewayArn, "/")
	rc.APIID = ss[0]
	rc.Stage = ss[1]
	rc.Method = ss[2]
	rc.Resource = ss[3]
	return rc
}
