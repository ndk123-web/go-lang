package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ndk123-web/muxrouter/dbconnection"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var client = dbconnection.ConnectMongoDB()

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Insertion In Mongodb (Create)
	collection := client.Database("golangdb").Collection("users")
	newUser := dbconnection.User{
		ID:   bson.NewObjectID(),
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	json.NewEncoder(w).Encode(insertResult)
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	// Read DB
	filter := bson.D{{"name", "Alice"}}
	var isUserExist dbconnection.User

	collection := client.Database("golangdb").Collection("users")

	err := collection.FindOne(context.TODO(), filter).Decode(&isUserExist)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document was found with the filter", filter)
			return
		} else {
			panic(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Found a single document: %+v\n", isUserExist)
	json.NewEncoder(w).Encode(isUserExist)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Update DB

	collection := client.Database("golangdb").Collection("users")

	filter := bson.D{{"name", "Alice"}}
	update := bson.D{{"$set", bson.D{{"age", 35}, {"city", "Boston"}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Matched %v document(s) and modified %v document(s)\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateResult)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	collection := client.Database("golangdb").Collection("users")

	filter := bson.D{{"name", "Alice"}}
	// Delete DB
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deleteResult)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filter := bson.D{{}}
	collection := client.Database("golangdb").Collection("users")

	allUsers, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Users: ", allUsers)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Get All Users Endpoint"})
}
