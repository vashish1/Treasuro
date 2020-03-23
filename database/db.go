package database

import (
	"context"
	"fmt"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Createdb creates a database
func Createdb() (*mongo.Collection,*mongo.Collection) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	usercollection := client.Database("Treasuro").Collection("User")
    uuid :=client.Database("Treasuro").Collection("Uid")
	return usercollection, uuid
}

