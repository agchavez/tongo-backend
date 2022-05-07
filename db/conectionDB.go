package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCon = connectDB()

var clientOption = options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
	AuthMechanism: "SCRAM-SHA-1",
	AuthSource:    "admin",
	Username:      "root",
	Password:      "root",
})

// ConnectDB is a function to connect to MongoDB
func connectDB() *mongo.Client {
	fmt.Println("asadsasd", os.Getenv("DB_USER"))
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func CheckConnection() int {
	err := MongoCon.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
