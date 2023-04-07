package main

import (
	"addison/helper"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func searchTrack(w http.ResponseWriter, r *http.Request) {
	var newTrack helper.Track

	if r.Method != "POST" {
	}

	log.Println(r.UserAgent())
	userAgent := r.UserAgent()
	if !strings.Contains(userAgent, "curl") {
		r.ParseForm()
		data := url.Values{
			"audio":     {r.FormValue("Audio")},
			"return":    {"spotify"},
			"api_token": {"50ed66b644ea194d65f31cd8a5a27778"},
		}

		response, err := http.PostForm("https://api.audd.io/", data)
		if err != nil {

		}

		defer response.Body.Close()

		body, _ := io.ReadAll(response.Body)

		var responseFromAuddio helper.Response
		if err := json.Unmarshal(body, &responseFromAuddio); err != nil {
			log.Println("HError222:", err)
			return
		}
		log.Println(responseFromAuddio.Result.Title)

		fmt.Fprintf(w, responseFromAuddio.Result.Title)
		return

	}

	////////////////////////////////////

	reqBody, err := ioutil.ReadAll(r.Body)

	//newTrack.Audio = audioStr
	//fmt.Println(reqBody)
	json.Unmarshal(reqBody, &newTrack)
	//fmt.Fprintf(w, newTrack.Audio)

	if err != nil {

	}

	data := url.Values{
		"audio":     {newTrack.Audio},
		"return":    {"spotify"},
		"api_token": {"50ed66b644ea194d65f31cd8a5a27778"},
	}

	//log.Println(data)
	response, err := http.PostForm("https://api.audd.io/", data)
	//log.Println(response.Body)
	if err != nil {
		//panic(err)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	// Parse the response into a Response struct
	var responseFromAuddio helper.Response
	if err := json.Unmarshal(body, &responseFromAuddio); err != nil {
		log.Println("Error:", err)
		return
	}
	resStr := "{\"Id\":" + responseFromAuddio.Result.Title + "\"}"
	fmt.Fprintf(w, resStr)

	if err != nil {

	}

}
func main() {

	http.HandleFunc("/search", searchTrack)
	fmt.Printf("Starting server at port 3001\n")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
