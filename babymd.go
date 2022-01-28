package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	in := `# Title

[a link!](https://something.com)

[NotYet](DoNotAddThisLink.com)

[another link!](some-page.html)
`
	err := md(strings.NewReader(in), os.Stdout)
	if err != nil {
		panic(err)
	}
}

// Markdown syntax: https://daringfireball.net/projects/markdown/syntax
func md(inIO io.Reader, outIO io.Writer) error {
	in := bufio.NewReader(inIO)
	out := bufio.NewWriter(outIO)
	defer out.Flush()
LINES:
	for {
		eol := []string{}
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break LINES
		}
		if err != nil {
			return err
		}
		// Short circuit for empty lines being newlines
		if len(line) == 1 && line[0] == '\n' {
			out.WriteString("<br>")
			continue LINES
		}
		// Since this will be a real line
		eol = append(eol, "<br>")

		// Headers
		if line[0] == '#' {
			n := headerN(line)
			out.WriteString(fmt.Sprintf("<h%d>", n))
			eol = append(eol, fmt.Sprintf("</h%d>", n))
			line = line[n:]
		}

		fmt.Println("Line:", line)
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			r := runes[i]
			switch r {
			case '[':
				// Link begin
				text, to, end, ok := getLink(line[i:])
				if !ok {
					// This is not a valid link, just write it out
					out.WriteRune(r)
				} else {
					// Valid link
					out.WriteString(fmt.Sprintf(`<a href="%s">%s</a>`, to, text))
					i = end
				}
			case '\n':
				// Do nothing
			default:
				out.WriteRune(r)
			}
		}

		// Do end of line closures
		for i := len(eol) - 1; i >= 0; i-- {
			out.WriteString(eol[i])
		}
	}
	return nil
}

// Get a link out of a body of text
func getLink(line string) (string, string, int, bool) {
	type State int
	const (
		Text State = iota
		Link
	)
	sq := 0
	par := 0
	text := ""
	to := ""
	state := Text
	ok := true
	end := 0

SCRAPE:
	for i, r := range line {
		switch {
		/* Text */
		case r == '[' && state == Text:
			sq++
			if sq > 1 {
				text += string(r)
			}

		/** Close **/

		case r == ']' && sq <= 1:
			// Text close
			state = Link
		case r == ']' && state == Text:
			sq--

		/* Links */

		case r == '(' && state == Link:
			par++
			if par > 1 {
				to += string(r)
			}

		/** Close **/

		case r == ')' && par <= 1:
			// Link close
			end = i
			break SCRAPE
		case r == ')' && state == Link:
			par--

		/* Default */

		default:
			switch state {
			case Text:
				text += string(r)
			case Link:
				to += string(r)
			}
		}
	}
	return text, to, end, ok
}

// Get the header depth of the line
func headerN(line string) (n int) {
	for _, r := range line {
		if r == '#' {
			n++
		} else {
			break
		}
	}
	return
}
