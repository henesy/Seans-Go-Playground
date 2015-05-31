package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
    "time"
//	"svi"
)

var scrn *[]Shape
var score, w, h int
var curShape *Shaper
var curPos int
var running bool
var hei int = 20
var wid int = 20
var dirTxt string
//var screen [23][10]occ

type key int

const (
	W key = iota
	A
	D
	S
	O
	P
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
	dir int
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
	dropDown()
	init()
	export() Shape
}

/* methods to satisfy Shaper() interface for type Shape
	0 = first (up)
	1 = second (right)
	2 = third (down)
	3 = fourth (left)
*/

/* rotates block left */
func (shp *s) rotateLeft() {

}

/* rotates block right */
func (shp *s) rotateRight() {

}

/* moves the block to the left */
func (shp *s) moveLeft() {
	// safety checks needed
	safe := true
	// 0, 3, 6 is border
	fut := (*shp)

	for i := 0; i < len(fut.blk); i++ {
		if fut.blk[i].x != 0 {
			fut.blk[i].x -= 1
		}
	}

	if (*shp).blk[0].x == 1 || (*shp).blk[3].x == 1 || (*shp).blk[6].x == 1 {
		safe = false
	}

	B1:
	for p, v := range (*scrn) {
		if p != curPos {
			for i := 0; i < len(v.blk); i++ {
				if (v.blk[i].x != 0 && v.blk[i].x == fut.blk[i].x) && (v.blk[i].y != 0 && v.blk[i].y == fut.blk[i].y) {
					safe = false
					break B1
				}
			}
		}
	}

	if safe == true {
		for i := 0; i < len(shp.blk); i++ {
			if (*shp).blk[i].x > 0 {
				(*shp).blk[i].x -= 1
			}
		}
	}
	dirTxt = "LEFT"
}

/* moves the block to the right */
func (shp *s) moveRight() {
	// safety checks needed
	safe := true
	fut := (*shp)


	if (*shp).blk[2].x == wid || (*shp).blk[5].x == wid || (*shp).blk[8].x == wid {
		safe = false
	}

	for i := 0; i < len(fut.blk); i++ {
		if fut.blk[i].x != 0 {
			fut.blk[i].x += 1
		}
	}

	if (*shp).blk[2].x == wid || (*shp).blk[5].x == wid || (*shp).blk[8].x == wid {
		safe = false
	}

	B1:
	for p, v := range (*scrn) {
		if p != curPos {
			for i := 0; i < len(v.blk); i++ {
				if (v.blk[i].x != 0 && v.blk[i].x == fut.blk[i].x) && (v.blk[i].y != 0 && v.blk[i].y == fut.blk[i].y) {
					safe = false
					break B1
				}
			}
		}
	}

	if safe == true {
		for i := 0; i < len(shp.blk); i++ {
			if (*shp).blk[i].x > 0 {
				(*shp).blk[i].x += 1
			}
		}
	}
	dirTxt = "RIGHT"
}

/* drops the block down at double the rate */
func (shp *s) dropDown() {

}

/* initializes the shape at the default '0' or 'UP' state */
func (shp *s) init() {
	(*shp).clr = termbox.ColorGreen
	(*shp).dir = 0

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

/* converts the current specific shape to a generic shape */
func (shp *s) export() Shape {
	return Shape(*shp)
}


/* picks a semi-random shape to provide the user next */
/*func randShape()(newShpr Shaper) {
	n := svi.Random(1, 7)

	if n == 1 {
		newShpr = Shaper(new(o))
	} else if n == 2 {
		newShpr = Shaper(new(i))
	} else if n == 3 {
		newShpr = Shaper(new(s))
	} else if n == 4 {
		newShpr = Shaper(new(z))
	} else if n == 5 {
		newShpr = Shaper(new(l))
	} else if n == 6 {
		newShpr = Shaper(new(j))
	}

	return
}*/

/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

/* handles input from the user via termbox-go */
func manageInput(screen []Shape, drawChan chan dir, runChan chan key) {
	for {
		select {
		case k := <- runChan:
			if k == D {
				(*curShape).moveRight()
			}
			if k == A {
				(*curShape).moveLeft()
			}
			if k == S {
				(*curShape).dropDown()
			}
			screen[curPos] = (*curShape).export()
		default:
		}
		<- drawChan
	}
}

/* draws and refreshes the screen */
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
		//x, y := 1, 1
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
				//x++
			}
			//y++
			//x = 1
		}

		tbPrint(25, 5, termbox.ColorWhite, termbox.ColorBlack, dirTxt)

		termbox.Flush()
        time.Sleep(60 * time.Millisecond)
		drawChan <- 0
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

	//array of blocks on the screen
	screen := make([]Shape, 1, 40)
	scrn = &screen
	/* init the interface as "s"  shape initially (rand() later) */
	shpr := Shaper(new(s))
	shpr.init()
	screen[0] = shpr.export()
	curPos = 0
	curShape = &shpr

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.Flush()
	go manageInput(screen, drawChan, runChan)
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
			} else if key == "o" {
				runChan <- O
			} else if key == "p" {
				runChan <- P
			}
		default:
		}
	}
}
