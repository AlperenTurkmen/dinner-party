package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tracks" {
		http.Error(w, "YOK BRO", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Tracks will be listed here.")
}

func main() {

	http.HandleFunc("/tracks", tracksHandler) // Update this line of code

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)

	}
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}
