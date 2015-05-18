package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
    "flag"
    //"math/rand"
    //"net"
)

/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
    for _, c := range msg {
            termbox.SetCell(x, y, c, fg, bg)
            x++
    }
}

func draw(w, h int) {

        for y := 0; y < h; y++ {
    		for x := 0; x < w; x++ {
    			tbPrint(x, y, termbox.ColorBlue, termbox.ColorBlack, "*")
    		}
    	}
}

/* client for SocketH rewritten in termbox-go */

func main() {
    words := ""
    flag.StringVar(&words, "a", "localhost:9090", "Set address to dial to for server.")
    flag.Parse()
    defer fmt.Print("\nGoodbye!\n")
    defer termbox.Close()

    err := termbox.Init()
    if err != nil {
        fmt.Println(err)
    }
    termbox.SetInputMode(termbox.InputAlt)
    w, h := termbox.Size()

    termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
    termbox.Flush()
    draw(w,h)
    termbox.Flush()
    pX, pY := 0, 0
    for run:=true;run==true; {
        switch ev := termbox.PollEvent(); ev.Type {
            case termbox.EventKey:
                key := string(ev.Ch)

                    if key == "q" {
                        pX++
                        pY++
                        tbPrint(pX, pY, termbox.ColorRed, termbox.ColorBlue, ":P")
                        termbox.Flush()
                    }
                    if key == "N" {
                        run=false
                        termbox.Flush()
                    }
                    draw(w, h)
        }
    }
}
