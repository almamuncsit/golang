// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"net/http"
	"context"
	"cloud.google.com/go/bigquery"
)


// Save data to cloud Storage
func saveToCloudStorage( w http.ResponseWriter, projectID string, datasetID string, tableID string ) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)

	newGcsRef := bigquery.NewGCSReference("gs://my-test-appsero-storage/marge_data/tracking.json")
	newGcsRef.DestinationFormat = bigquery.JSON
	
	ds := client.Dataset(datasetID)
	extractor := ds.Table(tableID).ExtractorTo(newGcsRef)
	extractor.DisableHeader = true
	
	new_job, err := extractor.Run(ctx)
	if err != nil {
		fmt.Fprintf(w, "Failed to run:")
	}
	new_status, err := new_job.Wait(ctx)
	if err != nil {
		fmt.Fprintf(w, "Failed on Wait:")
	}
	if new_status.Err() != nil {
		fmt.Fprintf(w, "Status error:")
	}
}



// Main functions
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	
	projectID := "mamun-appsero"
    datasetID := "tracking"
	tableID := "temp_sites"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		fmt.Fprintf(w, "bigquery.NewClient: %v", err)
	}
	defer client.Close()

	gcsRef := bigquery.NewGCSReference("gs://my-test-appsero-storage/track/*")
	gcsRef.SourceFormat = bigquery.JSON
	gcsRef.AutoDetect = true
	loader := client.Dataset(datasetID).Table(tableID).LoaderFrom(gcsRef)
	loader.WriteDisposition = bigquery.WriteTruncate

	job, err := loader.Run(ctx)
	if err != nil {
		fmt.Fprintf(w, "Failed to Run %v\n", err)
	}
	status, err := job.Wait(ctx)
	if err != nil {
		fmt.Fprintf(w, "Failed to Wait %v\n", err)
	}

	if status.Err() != nil {
		fmt.Fprintf(w, "job completed with error: %v", status.Err())
	}

	// Save as a json file
	saveToCloudStorage( w, projectID, datasetID, tableID )
	

	fmt.Fprintf(w, "Done\n")
	
}
