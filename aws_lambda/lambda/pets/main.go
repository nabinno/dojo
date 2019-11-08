package main

import (
	"context"
	"fmt"
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
	envAuthorizationHeaderName string = os.Getenv("AUTHORIZATION_HEADER_NAME")
	envUserPoolID              string = os.Getenv("USER_POOL_ID")
	envRoleNameOfAdmins        string = os.Getenv("ROLE_NAME_OF_ADMINS")
	envRoleNameOfUsers         string = os.Getenv("ROLE_NAME_OF_USERS")

	ginLambda  *ginadapter.GinLambda
	dynamodb   = dynamo.New(session.New(&aws.Config{Region: aws.String(envRegion)}))
	itemsTable = dynamodb.Table(os.Getenv("ITEMS_TABLE_NAME"))
	usersTable = dynamodb.Table(os.Getenv("USERS_TABLE_NAME")) // @todo 2019-10-31
)

// SimpleResponse is simple response struct
type SimpleResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Handler is the main entry point for Lambda. Receives a proxy request and returns a proxy response
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")

		r := gin.Default()
		r.Use(CORSMiddleware())

		// @todo 2019-11-03
		auth := r.Group("/", CognitoAuthorizationMiddleware())
		auth.GET("/pets", getPets)
		auth.GET("/pets/:id", getPet)
		auth.POST("/pets", createPet)
		// r.GET("/pets", getPets)
		// r.GET("/pets/:id", getPet)
		// r.POST("/pets", createPet)
		//
		// r.GET("/pets-lambda", getPets)
		// r.GET("/pets-none", getPets)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

// GET /pets
func getPets(c *gin.Context) {
	var pets = []Pet{}
	var limit int64 = 10
	if c.Query("limit") != "" {
		newLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 10
		} else {
			limit = int64(newLimit)
		}
	}
	if limit > 50 {
		limit = 50
	}

	username, ok := getCurrentUserName(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, SimpleResponse{Message: "Unauthorized!!!!"})
		return
	}

	err := itemsTable.
		Get("owner", username).
		Index("owner-index").
		Limit(limit).
		All(&pets)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			SimpleResponse{Message: fmt.Sprintf("Pets were not found by username (%v): %v", username, err)},
		)
		return
	}

	c.JSON(http.StatusAccepted, pets)
}

// GET /pets/:id
func getPet(c *gin.Context) {
	var petID = c.Param("id")
	var pet Pet

	err := itemsTable.
		Get("id", petID).
		One(&pet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, SimpleResponse{Message: fmt.Sprintf("Pet with id %v was not found", petID)})
		return
	}

	if username, ok := getCurrentUserName(c); ok && pet.Owner == username {
		// if the pet is owned by the user or they are an admin, return it.
		c.JSON(http.StatusAccepted, pet)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, SimpleResponse{Message: "Unauthorized"})
	}
}

// POST /pets
func createPet(c *gin.Context) {
	var newPet = Pet{}

	err := c.BindJSON(&newPet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, SimpleResponse{Message: fmt.Sprint("Params are not valid: %v", err)})
		return
	}

	newPet.ID = getUUID()
	newPet.Owner = c.Query("username")
	newPet.OwnerDisplayName = getCurrentUserDisplayNames(c)

	err = itemsTable.
		Put(newPet).
		Run()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, SimpleResponse{Message: fmt.Sprint("Does not save pet: %v", err)})
		return
	}

	c.JSON(http.StatusAccepted, newPet)
}
