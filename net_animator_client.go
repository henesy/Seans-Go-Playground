package main

import (
    "fmt"
    "net"
    "encoding/json"
tb  "github.com/nsf/termbox-go"
)


/* Checks and prints errors */
func check(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

/* prints to position (x,y) */
func tbPrint(x, y int, fg, bg tb.Attribute, msg string) {
	for _, c := range msg {
		tb.SetCell(x, y, c, fg, bg)
		x++
	}
}

/* draws the screen buffer */
func draw() {
	defer tb.Flush()
	for {


		tb.Flush()
	}
}

/* processes stream from server */
func recv() {


}


/* simply loops and prints the given world to the screen */
func main() {

    defer termbox.Flush()
    w, h := termbox.Size()
    err := termbox.Init()
    check(err)
    tb.SetInputMode(tb.InputAlt)
    err = tb.Clear(tb.Black, tb.White)
    check(err)
    tb.Flush()
    go draw()
    go recv()

	for {
		switch ev := tb.PollEvent(); ev.Type {
		case tb.EventKey:
			//key := string(ev.Ch)

			if ev.Key == tb.KeyCtrlQ {
	            break;
			}
		default:
		}
	}

}

