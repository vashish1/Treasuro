package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func InsertQuestion(c *mongo.Collection, u Question)bool {
	insertResult, err := c.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Print(err)
		return false
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

func UpdateScore(c *mongo.Collection, st string,sc int)bool {
	filter := bson.D{
		{"uuid", st},
	}
	update := bson.D{
		{
			"$inc", bson.D{{"score", sc},
		                   {"level",1}},
		},
		{
			"$set",bson.D{{"attempts",1}},
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

func UpdateAttempts(c *mongo.Collection,st string)bool{
	filter := bson.D{
		{"uuid", st},
	}
	update := bson.D{
		{
			"$inc", bson.D{{"attempts", 1},
		                   {"score",-2}},
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