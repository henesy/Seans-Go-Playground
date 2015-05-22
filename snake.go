package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	sc "strconv"
	"svi"
	"time"
	"unicode/utf8"
	"flag"
)

var w, h, score, length int = 0, 0, 0, 1

var difficulty uint64

type dir int

const (
	U dir = iota
	L
	R
	D
)

var tmpx string
var running, eaten bool = true, false
var sinceeat int

type sprite struct {
	r rune
	X int
	Y int
}

var target sprite

var snake = make([]sprite, 800)

func newTarget() {
	target.X, target.Y = svi.Random(0, w), svi.Random(1, h)
	for i := 0; i < length; i++ {
		if snake[i].X == target.X && snake[i].Y == target.Y && (target.X != 0 && target.Y != 0) && (target.X != w-1 && target.Y != h-1){
			newTarget()
		}
	}
}

/* has to check if player is on targets; repopulate targets */
func checkTarget() {
	if snake[0].X == target.X && snake[0].Y == target.Y {
		for i := len(snake) - 1; i > 0; i-- {
			if i-1 > 0 {
				snake[i] = snake[i-1]
			} else {
				snake[1].X = target.X
				snake[1].Y = target.Y
			}

		}

		//target = append(target[:i], target[i+1:]...)
		tmpx = sc.Itoa(snake[1].X) + sc.Itoa(snake[1].Y)
		length++
		score++
		eaten = true
		newTarget()

	} else {
		if eaten != true {
			eaten = false
		}
	}
}

func backTrace() {
	/* catch overlap and bug out */
	/* maybe change to just check impacts with [0] */
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if ((i-1 != j && j-1 != i && j+1 != i && i+1 != j) && i != j) && (snake[i].X == snake[j].X && snake[i].Y == snake[j].Y) {
				running = false
				i = length+1
				break
			}
		}
	}

	for i := length-1; i > 0; i-- {
            snake[i].X = snake[i-1].X
            snake[i].Y = snake[i-1].Y
    }
}

/* move the snake head; invokes shiftParts() */
func moveSnake(moveChan chan dir, drawChan chan dir, pauseChan chan bool) {
	var d dir
	for run := true;run == true; {
		select {
		case d = <- moveChan:
		case b := <- pauseChan:
			for b == true {
				b = <- pauseChan
			}

		default:
			//oldX, oldY := snake[0].X, snake[0].Y
		    if eaten == false {
		        backTrace()
		    } else {
		        eaten = false
		    }

		    if d == U && (snake[0].Y-1 > 0) {
				snake[0].Y--
			} else if d == D && (snake[0].Y+1 < h) {
				snake[0].Y++
			} else if d == L && (snake[0].X-1 > -1) {
				snake[0].X--
			} else if d == R && (snake[0].X+1 < w) {
				snake[0].X++
			}
		}
		drawChan <- d
	}
}

/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

/* draws the screen/updates */
func draw(w, h int, drawChan chan dir) {
	for {
		defer termbox.Flush()

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlue)
		for p, v := range snake {
			if p > length {
				break
			}
			tbPrint(v.X, v.Y, termbox.ColorRed, termbox.ColorBlue, "■")
			//tbPrint(35, 0, termbox.ColorWhite, termbox.ColorBlue, "Pos Snake: "+sc.Itoa(p))
		}
		scorestr := "Score: " + sc.Itoa(score)
		sssize := utf8.RuneCountInString(scorestr)
		tbPrint(1, 0, termbox.ColorWhite, termbox.ColorBlack, scorestr)
		for x := sssize + 1; x < w; x++ {
			tbPrint(x, 0, termbox.ColorBlack, termbox.ColorBlack, "█")
		}
		//fixes the printing of extras...which shouldn't happen, but w/e
		tbPrint(0, 0, termbox.ColorBlack, termbox.ColorBlack, "■")
		tbPrint(target.X, target.Y, termbox.ColorCyan, termbox.ColorBlue, string(target.r))
		checkTarget()

		termbox.Flush()
		d := <- drawChan
		if d == U || d == D {
			time.Sleep(60 * time.Millisecond)
		} else {
			time.Sleep(35 * time.Millisecond)
		}
	}
}

/* A basic snake game implemented in termbox-go and Golang */

func main() {
	flag.Uint64Var(&difficulty, "d", 1, "Set difficulty [1+]")
	flag.Parse()
	defer func() {
		termbox.Close()
		fmt.Print("Your score: ", score, "\n")
	}()
	moveChan := make(chan dir, 4)
	drawChan := make(chan dir, 1)
	pauseChan := make(chan bool, 1)
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
	}
	termbox.SetInputMode(termbox.InputAlt)
	w, h = termbox.Size()
	snake[0].X, snake[0].Y = w/2, h/2
	target.X, target.Y, target.r = svi.Random(0, w), svi.Random(1, h), ''
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlue)
	termbox.Flush()
	go draw(w, h, drawChan)
	termbox.Flush()
	moveChan <- R

	go moveSnake(moveChan, drawChan, pauseChan)

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
					pauseChan <- true
					switch ev := termbox.PollEvent(); ev.Type {
					case termbox.EventKey:
						if ev.Key == termbox.KeyEnter {
							pauseChan <- false
							r = false
							break
						}
					default:
					}
				}
			} else if key == "w" {
                if snake[0].Y-1 > 0 {
					moveChan <- U
                } else {
                    running = false
                }
			} else if key == "a" {
                if snake[0].X-1 > -1 {
					moveChan <- L
                } else {
                    running = false
                }
			} else if key == "d" {
                if snake[0].X+1 < w {
					moveChan <- R
                } else {
                    running = false
                }
			} else if key == "s" {
                if snake[0].Y+1 < h {
					moveChan <- D
                } else {
                    running = false
                }
			}
		default:
		}
	}
}
