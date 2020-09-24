// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Send data to maongodb database
func insertIntoMongoDb(w http.ResponseWriter, data map[string]string) {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://appsero:abc123456@cluster0.5wcxx.gcp.mongodb.net/appsero?retryWrites=true&w=majority"))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://34.87.16.169:27017"))
	if err != nil {
		fmt.Fprintf(w, "Failed to open client\n")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Fprintf(w, "Failed to open context\n")
	}
	defer client.Disconnect(ctx)

	collection := client.Database("appsero").Collection("sites")
	res, _ := collection.InsertOne(ctx, data)
	fmt.Fprintf(w, "%v", res)

}

// Save Json file to cloud storage
func saveJSONToCloudStorage(w http.ResponseWriter, data map[string]string) {
	bucket := "my-test-appsero-storage"
	object := "track/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Fprint(w, "storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	bkt := client.Bucket(bucket)
	obj := bkt.Object(object)
	wc := obj.NewWriter(ctx)

	wc.ContentType = "text/josn"
	mapData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Error %v\n", err.Error())
	}

	if _, err := wc.Write(mapData); err != nil {
		fmt.Fprintf(w, "Failed to write json file %v\n", err)
	}

	if err := wc.Close(); err != nil {
		fmt.Fprintf(w, "Failed to close file %v\n", err)
	}
}

//HelloWord is the main function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	formErr := r.ParseMultipartForm(32)
	if formErr != nil {
		fmt.Fprintf(w, "ParseForm() Err: %v", formErr)
	}

	data := make(map[string]string)
	for key, value := range r.Form {
		new_key := strings.ReplaceAll(strings.TrimRight(key, "]"), "[", "_")
		data[new_key] = value[0]
	}

	// Store Json file to cloud storage
	saveJSONToCloudStorage(w, data)

	// Save data to MongoDB
	insertIntoMongoDb(w, data)

}
