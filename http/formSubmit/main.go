package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func convertToJSON(w http.ResponseWriter, dataMap map[string]string) {
	mapData, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Fprintf(w, "Error %v\n", err.Error())
	}

	jsonStr := string(mapData)
	fmt.Fprintf(w, "%s", jsonStr)
	t := time.Now()
	fmt.Println(t.Format("2006/01/02/15"))
	// Create folder If doesn't exist
	if _, err := os.Stat("track/" + t.Format("2006/01/02")); os.IsNotExist(err) {
		os.Mkdir("track/"+t.Format("2006/01/02"), 0777)
	}

	_ = ioutil.WriteFile("track/"+t.Format("2006/01/02")+strconv.FormatInt(time.Now().UnixNano(), 10)+".json", mapData, 0644)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":

		// err := r.ParseMultipartForm(32)
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}

		data := make(map[string]string)
		for key, value := range r.Form {
			newKey := strings.ReplaceAll(strings.TrimRight(key, "]"), "[", "_")
			data[newKey] = value[0]
			// data[key] =
		}

		convertToJSON(w, data)

	}
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8282", nil); err != nil {
		log.Fatal(err)
	}
}
