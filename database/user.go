package database

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"log"
	"github.com/google/uuid"
)

//User ......
type User struct {
	Name         string
	Email        string
	PasswordHash string
	Token        string
}

//Newuser .....
func Newuser(name string, email string, password string, img string) User {

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

//GenerateUUID generates a unique id for every user.
func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}

//UpdateToken updates the user info
func UpdateToken(c *mongo.Collection, o string, t string) bool {
	filter := bson.D{
		{"email", o},
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
