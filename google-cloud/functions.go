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
)

// Save Json file to cloud storage
func saveJSONToCloudStorage(w http.ResponseWriter, data map[string]string) {
	bucket := "my-test-appsero-storage"
	t := time.Now()
	object := "track/" + t.Format("2006/01/02/15") + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".json"
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

// HelloWorld is the main function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	formErr := r.ParseMultipartForm(32)
	if formErr != nil {
		fmt.Fprintf(w, "ParseForm() Err: %v", formErr)
	}

	data := make(map[string]string)
	for key, value := range r.Form {
		newKey := strings.ReplaceAll(strings.TrimRight(key, "]"), "[", "_")
		data[newKey] = value[0]
	}

	// Store Json file to cloud storage
	saveJSONToCloudStorage(w, data)

}
