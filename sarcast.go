package main

import (
	"bufio"
	"os"
	"io"
	"unicode"
)

// Sarcastify text — WhAt dO YoU MeAn
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	U := true

	for {
		r, _, err := r.ReadRune()
		if err != nil {
			if err != io.EOF {
				w.Write([]byte("err: " + err.Error()))
			}
			break
		}
		
		if !unicode.IsSpace(r) {
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

