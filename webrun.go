package main

import (
	"strings"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"flag"
)

var (
	port	string	// Port to listen on
	defCmd	string	// Default cmd
)

// Handler for GET requests to /
// Format: /ls&-l&/ducks
func handler(w http.ResponseWriter, r *http.Request) {
	var cmd string
	var args, fields []string
	p := r.URL.Path

	if len(p) < 2 {
		fields = strings.Split(defCmd, "&")
		cmd = fields[0]
		args = fields[1:]
		goto docmd
	}

	//p = strings.Replace(p, " ", "\\ ", -1)
	fields = strings.Split(p[1:], "&")

	if len(fields) < 2 {
		args = []string{}
	} else {
		args = fields[1:]
	}

	cmd = fields[0]

	docmd:
	log.Println("Running:", cmd, args)
	
	run := exec.Command(cmd, args...)
	out, err := run.CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, string(out))
	}
}

// Runs programs based on url string in GET
func main() {
	flag.StringVar(&defCmd, "c", "whoami", "Default command")
	flag.StringVar(&port, "p", ":8080", "Port to listen on")
	flag.Parse()
	//args := flag.Args()
	
	http.HandleFunc("/", handler)

	log.Println("Listening: tcp!*!" + port[1:])

	log.Fatal(http.ListenAndServe(port, nil))
}

