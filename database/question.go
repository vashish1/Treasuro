package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vashish1/Treasuro/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Level  int    `json:"level,omitempty"`
	Answer string `json:"answer,omitempty"`
}

type Question struct {
	Id       int
	Question string
	Answer   string
}

type LeaderBoard struct {
	Username string `json:"username,omitempty"`
	Score    int    `json:"score,omitempty"`
}

func FindQuestion(c *mongo.Collection, id int) Question {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
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
			"$set", bson.D{ {"timestamp", t}},
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

func Leaderboard(c *mongo.Collection)[]LeaderBoard {
	filter := bson.D{}
	opts:=options.Find()
	opts.SetSort(bson.D{{"score", -1},
                       {"timestamp",-1}})
	projection := bson.D{
		{"username", 1},
		{"score", 1},
	}
	var result []LeaderBoard
	cur, err := c.Find(context.Background(), filter, opts.SetProjection(projection))

	if err != nil {
		fmt.Println("the error is:", err)
	}
	for cur.Next(context.TODO()) {
		var elem LeaderBoard
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

func CheckAnswer(c *mongo.Collection,q Response)(int){
	filter := bson.D{{
		"id",q.Level},
		{"answer",q.Answer},
	}
	var result Question
    
	err := c.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(err)
	fmt.Println(result)
	if err != nil {
		res:=FindQuestion(c,0)
		if res.Answer==q.Answer{
		rndmsc := utilities.RandomScore()
		fmt.Println("rn",rndmsc)
		return rndmsc
		}
		return 0
	}
    return 10
}

func FixAttempts(c *mongo.Collection, st string) bool {
	filter := bson.D{
		{"uuid", st},
	}
	update := bson.D{
		{
			"$set", bson.D{ {"attempts", 0}},	
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