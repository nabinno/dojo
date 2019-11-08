package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

// Pet is data struct
type Pet struct {
	ID               string    `dynamo:"id"` // Primary Partition Key (Hash Key)
	Breed            string    `dynamo:"breed"`
	Name             string    `dynamo:"name"`
	DateOfBirth      time.Time `dynamo:"dateOfBirth"`
	Owner            string    `dynamo:"owner"` // Global Secondary Index
	OwnerDisplayName string    `dynamo:"ownerDisplayName"`
}

func getUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return uuid.String()
}
