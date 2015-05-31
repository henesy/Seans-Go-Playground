package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
    "time"
)

var score, w, h int
var running bool
var hei int = 20
var wid int = 20
//var screen [23][10]occ

type key int

const (
	W key = iota
	A
	D
	S
)

type dir int

const (
	UP key = iota
	DN
	LE
	RI
)

type block struct {
    x int
    y int
}

type Shape struct {
	blk [9]block
	clr termbox.Attribute
	shp TyShape
}

type TyShape rune

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
	init()
	export() Shape
}

/* methods to satisfy Shaper() interface for type Shape */

/* rotates block left */
func (shp *s) rotateLeft() {

}

func (shp *s) rotateRight() {

}

func (shp *s) moveLeft() {

}

func (shp *s) moveRight() {

}

func (shp *s) init() {
	// top row
	(*shp).blk[0].x = 0
	(*shp).blk[0].y = 0
	(*shp).blk[1].x = 3
	(*shp).blk[1].y = 1
	(*shp).blk[2].x = 4
	(*shp).blk[2].y = 1
	// middle row
	(*shp).blk[3].x = 2
	(*shp).blk[3].y = 2
	(*shp).blk[4].x = 3
	(*shp).blk[4].y = 2
	(*shp).blk[5].x = 0
	(*shp).blk[5].y = 0
	// bottom row
	(*shp).blk[6].x = 0
	(*shp).blk[6].y = 0
	(*shp).blk[7].x = 0
	(*shp).blk[7].y = 0
	(*shp).blk[8].x = 0
	(*shp).blk[8].y = 0



}

func (shp *s) export() Shape {
	return Shape(*shp)
}


/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func draw(w, h int, drawChan chan dir, screen []Shape) {
	for {
		defer termbox.Flush()

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)


		/* draw box frames */
		/* corners for screen 20x20*/
		tbPrint(0, 0, termbox.ColorWhite, termbox.ColorBlack, "╔")
		tbPrint(0, hei+1, termbox.ColorWhite, termbox.ColorBlack, "╚")
		tbPrint(wid+1, 0, termbox.ColorWhite, termbox.ColorBlack, "╗")
		tbPrint(wid+1, hei+1, termbox.ColorWhite, termbox.ColorBlack, "╝")
		/* bars for screen */
		for y := 0; y < hei+2; y++ {
			for x := 0; x < wid+2; x++ {
				if (y == 0 || y == hei+1) && (x != 0 && x != wid+1) {
					tbPrint(x, y, termbox.ColorWhite, termbox.ColorBlack, "═")
				}
				if (x == 0 || x == wid+1) && (y != 0 && y != hei+1) {
					tbPrint(x, y, termbox.ColorWhite, termbox.ColorBlack, "║")
				}
			}
		}

		/* draw screen[][] */
		x, y := 1, 1
		for i := 0; i < hei; i++ {
			for j := 0; j < wid; j++ {
				/* print blocks */
				for k := 0; k < len(screen); k++ { //number of shapes
					for l := 0; l < len(screen[k].blk); l++ { //number of blocks
						if screen[k].blk[l].x > 0 && screen[k].blk[l].y > 0 {
							tbPrint(screen[k].blk[l].x, screen[k].blk[l].y, screen[k].clr, termbox.ColorBlack, "█")
						}
					}
				}

				x++
			}
			y++
			x = 1
		}

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
	runChan := make(chan key, 1)
	drawChan := make(chan dir, 1)

    err := termbox.Init()
	if err != nil {
		fmt.Println(err)
	}
	termbox.SetInputMode(termbox.InputAlt)
	w, h = termbox.Size()


	screen := make([]Shape, 1, 20)
	/* init the interface as "s"  shape initially (rand() later) */
	shpr := Shaper(new(s))
	shpr.init()
	screen[0] = shpr.export()


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
