package main

import (
	"bufio"
	"os"
	"io"
	"unicode"
	"flag"
)

var (
	tiny	bool	// make tiny text y/n
)

// Sarcastify text — WhAt dO YoU MeAn
func main() {
	flag.BoolVar(&tiny, "t", false, "Make tiny text instead")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

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

