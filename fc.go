package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"io"
	"log"
	"strconv"
)

var (
	nl string = "\n"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-lce] [-a n -b n | -p n] [file] n\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}


// fc - find character
// Read to the n'th character (indexed from 0) and print it using stdin or a file
func main() {
	flag.Usage = usage
	pl := flag.Bool("l", false, "Print the line number rather than the character")
	cm := flag.Bool("c", false, "Use char size of 1 byte rather than rune")
	bn := flag.Int("b", 0, "Number of characters before n to print")
	an := flag.Int("a", 0, "Number of characters after n to print")
	pn := flag.Int("p", 0, "Number of characters padding before/after n to print")
	uesc := flag.Bool("e", false, "Don't the escaped forms of characters")
	// TODO - maybe windows \r\n or \r mode for newlines?
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	argc := len(args)
	out := bufio.NewWriter(os.Stdout)
	var n int
	var in *os.File
	sf := "%q"

	/* Set up environment */

	// p ^ (b || a)
	if *pn > 0 && (*an > 0 || *bn > 0) {
		usage()
	}

	if *pn > 0 {
		*an = *pn
		*bn = *pn
	}

	if *uesc {
		sf = "%s"
	}

	// Need file && char#
	switch {
	case argc == 1:
		n = atoi(args[0])
		in = os.Stdin
	case argc == 2:
		n = atoi(args[1])

		f, err := os.Open(args[0])
		if err != nil {
			log.Fatal("err: could not open file - ", err)
		}
		defer f.Close()
		in = f
	default:
		usage()
	}

	/* Read through the characters */

	r := bufio.NewReader(in)

	// Note: character in this case is taken to mean a full rune by default
	var read func(*bufio.Reader) (string, error) = readRune

	if *cm {
		read = readByte
	}

	nlc := 0
	notpassed := true

	var i int
	for i = 0 ; ; i++ {
		s, err := read(r)
		if err != nil {
			if err != io.EOF {
				log.Fatal("err: read failed w/o EOF - ", err)
			}
			break
		}

		if s == nl && notpassed {
			nlc++
			if i >= n {
				notpassed = false
			}
		}

		// Sorry
		if !*pl {
			bmin := max(n - *bn, 0)
			amax := max(n + *an, n)

			if i >= bmin && i <= amax {
				out.WriteString(clean(fmt.Sprintf(sf, s)))
			}
		}
	}

	if n >= i {
		log.Fatalf("err: specified char #%d and max char was #%d (%d char total)", n, i-1, i)
	}

	if *pl {
		out.WriteString(fmt.Sprintf("%s:%d", in.Name(), nlc))
	}

	out.WriteRune('\n')
	out.Flush()
}

// Read and return a rune
func readRune(r *bufio.Reader) (string, error) {
	ru, _, err := r.ReadRune()
	return string(ru), err
}

// Read and return a byte
func readByte(r *bufio.Reader) (string, error) {
	b, err := r.ReadByte()
	return string(b), err
}

// Max of two ints -- if a == b, return a
func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

// We want \t, but not "\t"
// We want b, but not "b"
func clean(s string) string {
	if len(s) < 3 {
		return s
	}

	r := s[1:]

	return r[:len(r)-1]
}

// Wrap sc.Atoi()
func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("err: n is not an integer - ", err)
	}

	return n
}
