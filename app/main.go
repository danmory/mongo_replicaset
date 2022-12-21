package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://mongo1:27017,mongo2:27017/?replicaSet=rs1"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to MongoDB!")
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Ping connect error: %+v", err)
	}
	for i := 0; ; i++ {
		doc := bson.M{"name": i}
		if _, err := client.Database("test").Collection("abc").InsertOne(context.TODO(), doc); err != nil {
			log.Fatal(err)
		}
		log.Printf("Inserted %v", i)
		time.Sleep(5 * time.Second)
	}
}
