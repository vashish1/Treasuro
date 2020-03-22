package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Client) {

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
	usercollection := client.Database("Dev").Collection("User")
	return usercollection, client
}

//Insertintouserdb inserts the data into the database
func Insertintouserdb(usercollection *mongo.Collection, u User) bool {

	fmt.Println(u.Name)
	insertResult, err := usercollection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

//Findfromuserdb finds the required data
func Findfromuserdb(usercollection *mongo.Collection, st string, p string) bool {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if result.PasswordHash != SHA256ofstring(p) {
		return false
	}
	return true
}

//Finddb finds the required database
func Finddb(c *mongo.Collection, s string) User {
	filter := bson.D{primitive.E{Key: "email", Value: s}}
	var result User

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}
	return result
}
