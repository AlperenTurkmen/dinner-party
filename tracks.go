package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Track struct {
	Data string `json:"data"`
}

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/tracks/") {
		http.Error(w, "YOK BRO", http.StatusNotFound)
		return
	}

	if r.Method != "PUT" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the track ID from the URL
	parts := strings.Split(r.URL.Path, "/")
	trackID := parts[len(parts)-1]

	// Read the request body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body.", http.StatusBadRequest)
		return
	}

	// Encode the file to base64
	encodedData, err := encodeToBase64(bodyBytes)
	if err != nil {
		http.Error(w, "Error encoding file to Base64.", http.StatusInternalServerError)
		return
	}

	// Save the encoded data to a JSON file
	track := Track{
		Data: encodedData,
	}
	trackData, err := json.Marshal(track)
	if err != nil {
		http.Error(w, "Error encoding JSON data.", http.StatusInternalServerError)
		return
	}
	err = ioutil.WriteFile(trackID+".json", trackData, 0644)
	if err != nil {
		http.Error(w, "Error writing to file.", http.StatusInternalServerError)
		return
	}
	if len(bodyBytes) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Return a success message and set the HTTP status code to 201
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Track data saved to %s.json.", trackID)
}

func encodeToBase64(file []byte) (string, error) {
	encoded := base64.StdEncoding.EncodeToString(file)
	return encoded, nil
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file.", http.StatusBadRequest)
		return
	}

	encoded, err := encodeToBase64(fileBytes)
	if err != nil {
		http.Error(w, "Error encoding file to Base64.", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(encoded))
}

func main() {
	http.HandleFunc("/tracks/", tracksHandler)
	http.HandleFunc("/encode/", encodeHandler)

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
