package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Select the collection
	collection := client.Database("testing").Collection("numbers")

	// Find and display the result
	cursor, err := collection.Find(ctx, bson.M{"name": "pi"})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var numbers []bson.M
	if err = cursor.All(ctx, &numbers); err != nil {
		log.Fatal(err)
	}
	for _, number := range numbers {
		fmt.Println(number["name"])
	}

	// Run Another Query usnig cursor.Next
	fmt.Println(" === Another query ===")
	if cursor, err = collection.Find(ctx, bson.M{}); err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var number bson.M
		err := cursor.Decode(&number)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(number["name"])
	}

	// Find One Docuemnt
	fmt.Println(" === Find One Docuemnt ===")
	var singleNumber bson.M
	if err = collection.FindOne(ctx, bson.M{}).Decode(&singleNumber); err != nil {
		log.Fatal(err)
	}
	fmt.Println(singleNumber)

	// Sort Query Resutl
	fmt.Println(" === Find Sorted Docuemnt ===")
	opts := options.Find()
	opts.SetSort(bson.D{{"value", -1}})
	sortCursor, err := collection.Find(ctx, bson.D{{"value", bson.D{{"$gt", 4}}}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var numberSorted []bson.M
	if err = sortCursor.All(ctx, &numberSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(numberSorted)
}
