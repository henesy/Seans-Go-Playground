package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	playURL string = "https://play.golang.org"
)

type Output struct {
	Errors      string  `json: "Errors, omitempty"`
	Events      []Event `json: "Events, omitempty"`
	Status      int     `json: "Status, omitempty"`
	IsTest      bool    `json: "IsTest, omitempty"`
	TestsFailed int     `json: "TestsFailed, omitempty"`
	VetOK       bool    `json: "VetOK, omitempty"`
	ShareURL	string
	RawOut		string
}

type Event struct {
	Message string `json: "Message, omitempty"`
	Kind    string `json: "Kind, omitempty"`
	Delay   int    `json: "Delay, omitempty"`
}

// Send html-encoded input to play.golang.org and print the output, if any
func main() {
	raw := flag.Bool("r", false, "Print the raw JSON response instead")
	exit := flag.Bool("e", false, "Show the exit status even if 0")
	space := flag.Bool("n", false, "Add an extra newline between outputs")
	share := flag.Bool("s", false, "Print a share URL before output")
	nocomp := flag.Bool("C", false, "Don't compile")
	// TODO - flag for printing to stderr if stderr tagged output message
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	out := bufio.NewWriter(os.Stdout)
	inputs := []*bufio.Reader{bufio.NewReader(os.Stdin)}

	/* Convert arguments into inputs */
	if argc := len(args); argc > 0 {
		inputs = make([]*bufio.Reader, 0, argc)
		for _, name := range args {
			f, err := os.Open(name)
			if err != nil {
				log.Fatal("err: could not open file - ", err)
			}
			defer f.Close()

			r := bufio.NewReader(f)
			inputs = append(inputs, r)
		}
	}

	/* Send inputs; read responses */
	sinputs := read2s(inputs)
	outputs := make([]Output, 0, len(inputs))

	for i, s := range sinputs {
		o := Output{}

		vals := url.Values{
			"version": {"2"},
			"body":    {s},
			"withvet": {"true"},
		}

		if *share {
			o.ShareURL = getshare(s) + "\n"
		}

		resp, err := http.PostForm(playURL+"/compile", vals)
		if err != nil {
			log.Fatal("err: POST req failed for input #", i, " - ", err)
		}

		/* Process response */

		// Read into buf as well
		buf := new(bytes.Buffer)

		buf.ReadFrom(resp.Body)

		if *nocomp {
			buf.Truncate(0)
		}

		resp.Body.Close()

		if *raw {
			o.RawOut = buf.String()
			buf.Truncate(0)
		}

		// Response is a JSON object
		d := json.NewDecoder(buf)

		err = d.Decode(&o)
		if err != nil {
			if err != io.EOF {
				log.Fatal("err: JSON decode of response failed for input #", i, " - ", err)
			}
		}

		outputs = append(outputs, o)
	}

	/* Print outputs */
	for _, o := range outputs {
		s := ""

		s += o.ShareURL + o.RawOut

		if o.Errors != "" {
			s = o.Errors
		} else {
			for _, e := range o.Events {
				s += e.Message
			}
		}

		out.WriteString(s)

		if *exit || o.Status != 0 {
			out.WriteString(fmt.Sprintf("\nProgram exited: status %d.\n", o.Status))
		}

		if *space {
			out.WriteRune('\n')
		}

		out.Flush()
	}

	out.Flush()
}

// Convert read out the strings in the readers
func read2s(inputs []*bufio.Reader) (runes []string) {
	runes = make([]string, 0, len(inputs))

	// Read input(s) and emit their encoded form
	for i, in := range inputs {
		if in == nil {
			log.Fatal("Got nil buffer on writer #", i)
		}
		s := ""

		for {
			r, _, err := in.ReadRune()
			if err != nil {
				if err != io.EOF {
					log.Fatal("err: read failed - ", err)
				}
				break
			}
			s += string(r)
		}

		runes = append(runes, s)
	}

	return
}

// Returns the share URL for a given string
func getshare(body string) string {
	r := strings.NewReader(body)

	resp, err := http.Post(playURL+"/share", "text/plain; charset=utf-8", r)
	if err != nil {
		log.Fatal("err: POST share failed - ", err)
	}

	buf := new(bytes.Buffer)

	buf.ReadFrom(resp.Body)

	resp.Body.Close()

	tail := buf.String()

	return playURL + "/p/" + tail
}

