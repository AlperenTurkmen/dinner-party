package main

import (
	"addison/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var allTracks helper.Tracks

func tracksLister(w http.ResponseWriter, r *http.Request) {

	//find track ID from URL

	parts := strings.Split(r.URL.Path, "/")

	trackID := parts[len(parts)-1]
	//w.WriteHeader(http.StatusCreated)
	// Read existing tracks from the file
	allTracks, err := ReadTracksFromFile("tracks.json")
	log.Println(trackID + "Geciyor")
	//return
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error reading tracks from file.", http.StatusInternalServerError)
		return

	}

	log.Println("trackID=" + trackID)

	switch r.Method {

	case "GET":
		if trackID == "tracks" || trackID == "" {
			var s string = "["

			for _, item := range allTracks {

				s = s + "\"" + item.ID + "\","
			}
			if len(s) > 2 {
				s = strings.TrimSuffix(s, ",")
			}
			s = s + "]"
			fmt.Fprintf(w, s)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			// Show track from ID
			found := false
			for _, singleTrack := range allTracks {
				if singleTrack.ID == trackID {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(singleTrack)

					//w.Write([]byte(singleTrack.Audio))
					if err != nil {

					}

					found = true
					break
				}
			}
			if !found {
				//	fmt.Fprintf(w, "not found needs to be 204")
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
	case "PUT":
		var newTrack helper.Track
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Kindly enter data with the track ID and audio only in order to update", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(reqBody, &newTrack)
		if err != nil {
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		if newTrack.ID == "" || newTrack.Audio == "" {
			http.Error(w, "Missing track ID or audio in input", http.StatusBadRequest)
			return
		}
		allTracks = append(allTracks, newTrack)
		content, err := json.Marshal(allTracks)
		if err != nil {
			http.Error(w, "Error marshaling JSON data", http.StatusInternalServerError)
			return
		}
		err = ioutil.WriteFile("tracks.json", content, 0644)
		if err != nil {
			http.Error(w, "Error writing tracks to file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		//json.NewEncoder(w).Encode(newTrack)
		return

	case "DELETE":
		var newAllTracks helper.Tracks
		log.Println("trackID:" + trackID)
		for _, singleTrack := range allTracks {
			if singleTrack.ID == trackID {
				continue
			}
			newAllTracks = append(newAllTracks, singleTrack)
		}
		content, err := json.Marshal(newAllTracks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = ioutil.WriteFile("tracks.json", content, 0644)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(newAllTracks) == len(allTracks) {
			// Track not found
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}

}

func ReadTracksFromFile(filename string) ([]helper.Track, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tracks []helper.Track
	err = json.Unmarshal(data, &tracks)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func main() {

	http.HandleFunc("/tracks/", tracksLister)
	http.HandleFunc("/tracks", tracksLister)
	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
