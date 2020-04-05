package database

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//User ......
type User struct {
	UUID         string
	Name         string
	Username     string
	PhnNo        string
	SQR          int
	Level        int
	Score        int
	Attempts     int
	TimeStamp    time.Time
	Email        string
	PasswordHash string
	Token        string
}

//Newuser .....
func Newuser(name, email, password, phn, sqr string, t time.Time) User {

	Password := SHA256ofstring(password)
	U := User{Name: name, Email: email, PasswordHash: Password}
	return U
}

//SHA256ofstring is a function which takes a string a returns its sha256 hashed form
func SHA256ofstring(p string) string {
	h := sha1.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
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
func Findfromuserdb(usercollection *mongo.Collection, st string) bool {
	filter := bson.D{primitive.E{Key: "uuid", Value: st}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CheckUsername(c *mongo.Collection,st string)bool{
	filter := bson.D{primitive.E{Key: "username", Value: st}}
	var result User

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

//FindUser finds if the user exists but with respect to the username.
func FindUser(usercollection *mongo.Collection, st string, p string) bool {
	filter := bson.D{primitive.E{Key: "username", Value: st}}
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
	filter := bson.D{primitive.E{Key: "uuid", Value: s}}
	var result User

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}
	return result
}

func UpdateUserCreds(c *mongo.Collection, id, username, phn, email, pass string) bool {
	filter := bson.D{
		{"uuid", id},
	}
	passhash := SHA256ofstring(pass)
	update := bson.D{
		{
			"$set", bson.D{
				{"email", email},
				{"passwordhash", passhash},
				{"username", username},
				{"phnno", phn},
				{"sqr", 5},
			    {"level",1}},
		},
	}
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}

//UpdateToken updates the user info
func UpdateToken(c *mongo.Collection, o string, t string) bool {
	filter := bson.D{
		{"uuid", o},
	}
	update := bson.D{
		{
			"$set", bson.D{{"token", t}},
		},
	}
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}

