// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"encoding/json"
	"net/http"
	"context"
	"time"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
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

	// Start code for form data
	formErr := r.ParseMultipartForm(32)
	if formErr != nil {
		fmt.Fprintf(w, "ParseForm() Err: %v", formErr)
	}

	data := make(map[string]string)
	for key, value := range r.Form {
		new_key := strings.ReplaceAll( strings.TrimRight(key, "]"), "[", "_" )
		data[new_key] = value[0]
	}
	// Finish storing data to Map 

	bkt := client.Bucket(bucket)
	obj := bkt.Object(object)
	wc := obj.NewWriter(ctx)


	// Write jon file
	wc.ContentType = "text/josn"
	mapData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Error %v\n", err.Error())
	}

	if _, err := wc.Write(mapData); err != nil {
        fmt.Fprintf(w, "Failed to write json file %v\n", err)
    }

	// Close, just like writing a file.
	if err := wc.Close(); err != nil {
	    fmt.Fprintf(w, "Failed to close file %v\n", err)
	}

}
