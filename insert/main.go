package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	StudentID int              `bson:"studentid"`
	Name      string           `bson:"name"`
	Age       int              `bson:"age"`
	Gender    string           `bson:"gender"`
	Grade     string           `bson:"grade"`
	Marks     []map[string]int `bson:"marks"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Environment variable MONGO_URI has to be set")
	}
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	newStudent := Student{
		StudentID: 1,
		Name:      "Varun CHandra",
		Age:       20,
		Gender:    "Male",
		Grade:     "A",
		Marks: []map[string]int{
			{"eng": 87},
			{"maths": 97},
			{"lang": 100},
		},
	}
	collection := client.Database("University").Collection("students")
	result, err := collection.InsertOne(context.TODO(), newStudent)
	if err != nil {
		panic(err)
	}
	fmt.Println("Document inserted with id", result.InsertedID)
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
