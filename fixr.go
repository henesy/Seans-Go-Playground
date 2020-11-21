package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"strings"
)

var (
	r	string
	to	string
	d	string
)

// Rename files wrapped in ' ' to not have those
func main() {
	flag.StringVar(&r, "r", "'", "Rune to replace in file name")
	flag.StringVar(&to, "t", "", "What to insert instead of rune")
	flag.StringVar(&d, "d", "./", "Directory to run the fix in")
	flag.Parse()

	if len(r) < 1 {
		log.Fatal("err: r cannot be empty")
	}

	f, err := os.Open(d)
	if err != nil {
		log.Fatal("err: could not open dir - ", err)
	}

	names, err := f.Readdirnames(-1)
	if err != nil {
		log.Fatal("err: could not read directory contents - ", err)
	}

	for _, n := range names {
		newn := strings.ReplaceAll(n, string(r[0]), to)
		fmt.Println("Moving", n, "to", newn)
		err := os.Rename(n, newn)
		if err != nil {
			log.Fatal("err: could not mv file ", n, " to ", newn, " - ", err)
		}
	}

	fmt.Println("Done.")
}

