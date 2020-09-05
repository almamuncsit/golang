package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func convertToJSON(w http.ResponseWriter, dataMap map[string]string) {
	mapData, err := json.MarshalIndent(dataMap, "", "")
	if err != nil {
		fmt.Fprintf(w, "Error %v\n", err.Error())
	}

	jsonStr := string(mapData)
	fmt.Fprintf(w, "%s", jsonStr)

	_ = ioutil.WriteFile(dataMap["name"]+".json", mapData, 0644)
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
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}

		data := make(map[string]string)
		for key, value := range r.Form {
			data[key] = value[0]
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