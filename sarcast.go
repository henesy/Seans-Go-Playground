package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"unicode"
	"flag"
	"strings"
	"github.com/atotto/clipboard"
)

var (
	tiny	bool	// make tiny text y/n
	clip	bool	// copy to clipboard y/n
)

// Output string hack for StringWriter
var outStr	string

// For use with clipboard functionality
type StringWriter struct {
	// Look up ☺
}

// Implement io.Writer
func (sw StringWriter) Write(p []byte) (n int, err error) {
	outStr += string(p)
	
	return len(p), nil
}

// Copy to clipboard
func copy() {
	err := clipboard.WriteAll(outStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to copy to clipboard --", err)
		os.Exit(1)
	}
}

// Sarcastify text — WhAt dO YoU MeAn
func main() {
	flag.BoolVar(&tiny, "t", false, "Make tiny text instead")
	flag.BoolVar(&clip, "c", false, "Copy to clipboard")
	flag.Parse()
	args := flag.Args()

	var from io.Reader
	var to io.Writer

	// Set input reader
	if len(args) > 0 {
		var str string
		for _, s := range args {
			str += s + " "
		}
		from = strings.NewReader(str + "\n")
	} else {
		from = os.Stdin
	}

	// Set output writer
	if clip {
		var b StringWriter
		to = b
		defer copy()
	} else {
		to = os.Stdout
	}
	
	r := bufio.NewReader(from)

	w := bufio.NewWriter(to)

	if tiny {
		tinify(r, w)
		return
	}

	// Sarcastic text
	U := true
	for {
		r, _, err := r.ReadRune()
		if err != nil {
			if err != io.EOF {
				w.Write([]byte("err: " + err.Error()))
			}
			break
		}
		
		if !(unicode.IsSpace(r) || unicode.IsPunct(r)) {
			if U {
				r = unicode.ToUpper(r)
				U = false
			} else {
				r = unicode.ToLower(r)
				U = true
			}
		}

		w.WriteRune(r)
	}
	w.Flush()
}

// Make text tiny -- only replace [a-z0-9]
func tinify(r *bufio.Reader, w *bufio.Writer) {
	smol := []rune{'ᵃ','ᵇ','ᶜ','ᵈ','ᵉ','ᶠ','ᵍ','ʰ','ᶦ','ʲ','ᵏ','ˡ','ᵐ','ⁿ','ᵒ','ᵖ','ᵠ','ʳ','ˢ','ᵗ','ᵘ','ᵛ','ʷ','ˣ','ʸ','ᶻ'}
	nums := []rune{'¹','²','³','⁴','⁵','⁶','⁷','⁸','⁹','⁰'}

	for {
		r, _, err := r.ReadRune()
		r = unicode.ToLower(r)
		
		if err != nil {
			if err != io.EOF {
				w.Write([]byte("err: " + err.Error()))
			}
			break
		}

		// Replace with small letter if able
		if unicode.IsLetter(r) {
			r = smol[r-'a']
		}

		if unicode.IsNumber(r) {
			r = nums[r-'0']
		}

		w.WriteRune(r)
	}
	w.Flush()
}

