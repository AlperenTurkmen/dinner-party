package main

import (
	"addison/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func coolTownApp(w http.ResponseWriter, r *http.Request) {

	var targetTrack helper.Track
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &targetTrack)

	data := url.Values{}
	data.Add("Audio", targetTrack.Audio)

	response, err := http.PostForm("http://localhost:3001/search", data)

	if err != nil {
		log.Println("PostForm error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ReadAll error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url := "http://localhost:3000/tracks/" + replacePlusWithSpace(string(body))

	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	} else if res.StatusCode == http.StatusNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/cooltown", coolTownApp)

	fmt.Printf("Starting server at port 3002\n")
	if err := http.ListenAndServe(":3002", nil); err != nil {
		log.Fatal(err)
	}
}

func replacePlusWithSpace(s string) string {
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	s = strings.ReplaceAll(s, " ", "+")
	return s
}
