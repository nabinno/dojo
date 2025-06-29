package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cIDP "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func getCurrentUserDisplayNames(c *gin.Context) string {
	firstName, ok1 := getCurrentUserFirstName(c)
	lastName, ok2 := getCurrentUserLastName(c)

	if !ok1 || !ok2 {
		idP := cIDP.New(session.New(&aws.Config{Region: aws.String(envRegion)}))
		username := c.Query("username")
		user, err := idP.AdminGetUser(&cIDP.AdminGetUserInput{Username: &username, UserPoolId: &envUserPoolID})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, SimpleResponse{Message: "User was not found"})
		}

		for _, userAttribute := range user.UserAttributes {
			if *userAttribute.Name == "name" || *userAttribute.Name == "given_name" {
				firstName = userAttribute.Value
			}
			if *userAttribute.Name == "family_name" {
				lastName = userAttribute.Value
			}
		}
	}

	return fmt.Sprintf("%v %v", firstName, lastName)
}

func getCurrentUserName(c *gin.Context) (username interface{}, ok bool) {
	username, ok = getClaims(c)["cognito:username"]
	if !ok {
		username = c.Query("username")
		if len(fmt.Sprintf("%v", username)) > 0 {
			ok = true
		}
	}
	return
}

func getCurrentUserFirstName(c *gin.Context) (firstName interface{}, ok bool) {
	firstName, ok = getClaims(c)["cognito:given_name"]
	return
}

func getCurrentUserLastName(c *gin.Context) (lastName interface{}, ok bool) {
	lastName, ok = getClaims(c)["cognito:family_name"]
	return
}

func getClaims(c *gin.Context) jwt.MapClaims {
	return c.MustGet("token").(*jwt.Token).Claims.(jwt.MapClaims)
}
