package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"time"
)

var (
	port	string
	count	uint64	= 0
)


// On GET, display counter, on POST, tick 1 up
func countHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Write([]byte(fmt.Sprint(count)))
	case "POST":
		count++
	}
}

// Writes the current time back
func timeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprint(time.Now())))
}


// Smol web app for Xylo
func main() {
	flag.StringVar(&port, "p", ":1337", "port to listen on")
	flag.Parse()

	http.HandleFunc("/count/", countHandler)
	http.HandleFunc("/time/", timeHandler)

	log.Print("Listening on http://localhost", port)

	log.Print(http.ListenAndServe(port, nil))
}

