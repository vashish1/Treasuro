package database

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UUID struct {
	UUID string
}

//GenerateUUIDs generates a unique id for every user.
func GenerateUUIDs() []string {
	var slice []string
	for i := 0; i < 50; i++ {
		sd := uuid.New()
		slice = append(slice, (sd.String()))
	}
	return slice
}

func InsertUUIDindb(c *mongo.Collection) bool {
	slice := GenerateUUIDs()
	for i := 0; i < 50; i++ {
		var id UUID
		id.UUID = slice[i]
		insertResult, err := c.InsertOne(context.TODO(), id)
		if err != nil {
			log.Print(err)
			return false
		}

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}
	return true
}

func UuidExists(c *mongo.Collection, u string) bool {
	var id UUID
	filter := bson.D{primitive.E{Key: "uuid", Value: u}}

	err := c.FindOne(context.TODO(), filter).Decode(&id)
	if err != nil {
		return false
	}
	return true
}
