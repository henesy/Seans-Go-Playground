package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
    "time"
)

var score, w, h int
var running bool
//var screen [23][10]occ

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

type Shape [9]block

type o Shape
type i Shape
type s Shape
type z Shape
type l Shape
type j Shape

type Shaper interface {
	rotateLeft()
	rotateRight()
	moveLeft()
	moveRight()
	initShape()
}


/* sets all of screen to unoccupied */
func unOccupy(screen [][]occ) {
	for i := 0; i < len(screen); i++ {
		for j := 0; j < len(screen[i]); j++ {
			screen[i][j] = UN
		}
	}
}

/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func draw(w, h int, drawChan chan dir, screen [][]occ) {
	for {
		defer termbox.Flush()

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)


		/* draw box frames */
		/* corners for screen */
		tbPrint(0, 0, termbox.ColorWhite, termbox.ColorBlack, "╔")
		tbPrint(0, len(screen)+1, termbox.ColorWhite, termbox.ColorBlack, "╚")
		tbPrint(len(screen[0])+1, 0, termbox.ColorWhite, termbox.ColorBlack, "╗")
		tbPrint(len(screen[0])+1, len(screen)+1, termbox.ColorWhite, termbox.ColorBlack, "╝")
		/* bars for screen */
		for y := 0; y < len(screen)+2; y++ {
			for x := 0; x < len(screen[0])+2; x++ {
				if (y == 0 || y == len(screen)+1) && (x != 0 && x != len(screen[0])+1) {
					tbPrint(x, y, termbox.ColorWhite, termbox.ColorBlack, "═")
				}
				if (x == 0 || x == len(screen[0])+1) && (y != 0 && y != len(screen)+1) {
					tbPrint(x, y, termbox.ColorWhite, termbox.ColorBlack, "║")
				}
			}
		}

		/* draw screen[][] */
		x, y := 1, 1
		for i := 0; i < len(screen); i++ {
			for j := 0; j < len(screen[i]); j++ {
				tbPrint(x, y, termbox.ColorBlue, termbox.ColorBlack, "*")
				x++
			}
			y++
			x = 1
		}
		/* draw occupancy in screen[][] */
		//screen[0][0] = OC
		//screen[len(screen)-1][len(screen[0])-1] = OC
		for y := 0; y < len(screen); y++ {
			for x := 0; x < len(screen[y]); x++ {
				if screen[y][x] == OC {
					tbPrint(x+1, y+1, termbox.ColorRed, termbox.ColorBlack, "█")
				}
			}
		}

		/* draw next block and dashboard */
		


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

	//22 tall, 20 wide, top 2 will be hidden
	var screen [][]occ = make([][]occ, 22)
	for i := 0; i < len(screen); i++ {
		screen[i] = make([]occ, 20)
	}
	unOccupy(screen)
	//hiddenScreen := screen[0:2]


	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.Flush()
	go draw(w, h, drawChan, screen)
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
