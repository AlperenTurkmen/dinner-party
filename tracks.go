package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Track struct {
	ID    string `json:"Id"`
	Audio string `json:"Audio"`
}
type Tracks []Track

func tracksLister(w http.ResponseWriter, r *http.Request) {
	var allTracks Tracks

	// Read existing tracks from the file
	allTracks, err := readTracksFromFile("tracks.json")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error reading tracks from file.", http.StatusInternalServerError)
		return
	}
	//find track ID from URL
	parts := strings.Split(r.URL.Path, "/")
	trackID := parts[len(parts)-1]
	fmt.Fprintf(w, "TRACK ID=")
	fmt.Fprintf(w, trackID)

	//Show track from ID
	for _, singleTrack := range allTracks {
		if singleTrack.ID == trackID {
			json.NewEncoder(w).Encode(singleTrack.Audio)
		}
	}
	//Return all tracks
	jsonBytes, err := json.Marshal(allTracks)
	if err != nil {
		// handle error
	}
	fmt.Fprintf(w, string(jsonBytes))

	return

}
func readTracksFromFile(filename string) ([]Track, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tracks []Track
	err = json.Unmarshal(data, &tracks)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func main() {

	http.HandleFunc("/tracks/", tracksLister)

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
