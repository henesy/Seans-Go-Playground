package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
    "time"
)

var score, w, h int
var running bool
//var screen [][]occ

type dir int

const (
	W dir = iota
	A
	D
	S
)

type occ int

const (
    UN occ = iota
    OC
)

type block struct {
    x int
    y int
}

type shape []block = make([]block, 9)

type o shape
type i shape
type s shape
type z shape
type l shape
type j shape

type Shaper interface {

}


/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func draw(w, h int, drawChan chan dir) {
	for {
		defer termbox.Flush()

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

		/* draw blocks @ current position */

		termbox.Flush()
        time.Sleep(60 * time.Millisecond)
	}
}


/* small termbox tetris game */

func main() {
    defer func() {
		termbox.Close()
		fmt.Print("Your score: ", score, "\n")
	}()
	runChan := make(chan dir, 1)
	drawChan := make(chan dir, 1)

    err := termbox.Init()
	if err != nil {
		fmt.Println(err)
	}
	termbox.SetInputMode(termbox.InputAlt)
	w, h = termbox.Size()
    screen := [h][w]occ

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.Flush()
	go draw(w, h, drawChan)
	termbox.Flush()
	runChan <- S

	for running = true; running == true; {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			key := string(ev.Ch)

			if ev.Key == termbox.KeyCtrlQ {
				running = false
				termbox.Flush()
			} else if ev.Key == termbox.KeyEnter {
				/* pause button */
                for r := true; r == true; {
					switch ev := termbox.PollEvent(); ev.Type {
					case termbox.EventKey:
						if ev.Key == termbox.KeyEnter {
							r = false
							break
						}
					default:
					}
				}
			} else if key == "w" {
					runChan <- W
			} else if key == "a" {
					runChan <- A
			} else if key == "d" {
					runChan <- D
			} else if key == "s" {
					runChan <- S
			}
		default:
		}
	}
}
