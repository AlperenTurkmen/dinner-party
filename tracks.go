package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Track struct {
	Data string `json:"data"`
}

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tracks" {
		http.Error(w, "YOK BRO", http.StatusNotFound)
		return
	}

	if r.Method != "PUT" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body.", http.StatusBadRequest)
		return
	}

	// Base64 decode the data
	decodedData, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		http.Error(w, "Error decoding Base64 data.", http.StatusBadRequest)
		return
	}

	// Save the decoded data to a JSON file
	track := Track{
		Data: string(decodedData),
	}
	trackData, err := json.Marshal(track)
	if err != nil {
		http.Error(w, "Error encoding JSON data.", http.StatusInternalServerError)
		return
	}
	err = ioutil.WriteFile("track.json", trackData, 0644)
	if err != nil {
		http.Error(w, "Error writing to file.", http.StatusInternalServerError)
		return
	}

	// Return a success message
	fmt.Fprintf(w, "Track data saved to track.json.")
}

func main() {
	http.HandleFunc("/tracks", tracksHandler)

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
