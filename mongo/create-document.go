package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Insert data into document
	data := make(map[string]string)
	data["name"] = "Hello"
	data["value"] = "12341234"

	collection := client.Database("testing").Collection("numbers")
	res, _ := collection.InsertOne(ctx, data)
	// res, _ := collection.InsertOne(ctx,
	// 	bson.M{
	// 		"name":     "pi",
	// 		"value":    3.14159,
	// 		"comments": bson.A{"Good", "Nice", "Bad"}})
	fmt.Println(res)

	// Insert Many Documents
	// manyRes, _ := collection.InsertMany(ctx, []interface{}{
	// 	bson.M{
	// 		"name":     "mamun",
	// 		"value":    89023845,
	// 		"comments": bson.A{"Good", "Bad"},
	// 	},
	// 	bson.M{
	// 		"name":     "Khan",
	// 		"value":    23532,
	// 		"comments": bson.A{"Good"},
	// 	},
	// })
	// fmt.Println(manyRes)

}
