package dbconnection

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
)

type User struct {
	ID   bson.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string        `bson:"name,omitempty" json:"name,omitempty"`
	Age  int           `bson:"age,omitempty" json:"age,omitempty"`
	City string        `bson:"city,omitempty" json:"city,omitempty"`
}

func ConnectMongoDB() *mongo.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	// Uses the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Creates a new client and connects to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// Sends a ping to confirm a successful connection

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
