package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Question struct {
	Id       int
	Question string
	Answer   string
}

func FindQuestion(c *mongo.Collection, id int) Question {
	filter := bson.D{primitive.E{Key: "Id", Value: id}}
	var result Question

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}
	return result
}

func InsertQuestion(c *mongo.Collection, u Question) bool {
	insertResult, err := c.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Print(err)
		return false
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

func UpdateScore(c *mongo.Collection, st string, sc int) bool {
	filter := bson.D{
		{"uuid", st},
	}
	t := time.Now()
	update := bson.D{
		{
			"$inc", bson.D{{"score", sc},
				{"level", 1}},
		},
		{
			"$set", bson.D{{"attempts", 1}, {"timestamp", t}},
		},
	}
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}

func UpdateAttempts(c *mongo.Collection, st string) bool {
	filter := bson.D{
		{"uuid", st},
	}
	update := bson.D{
		{
			"$inc", bson.D{{"attempts", 1},
				{"score", -2}},
		},
	}
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}

func Leaderboard(c *mongo.Collection)[]User {
	filter := bson.D{}

	projection := bson.D{
		{"_id", 0},
		{"username", 1},
		{"score", bson.D{
			{"$sort", -1}}},
		{"timestamp", bson.D{{
			"$sort", 1}}},
	}
	var result []User
	cur, err := c.Find(context.Background(), filter, options.Find().SetProjection(projection))

	if err != nil {
		fmt.Println("the error is:", err)
	}
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("decoding error:", err)

		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("cursor error", err)
	}

	cur.Close(context.TODO())
	fmt.Println(result)
	return result
}
