package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
    "flag"
    //"net"
)

var msgs []string = make([]string, 50)

/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
    for _, c := range msg {
            termbox.SetCell(x, y, c, fg, bg)
            x++
    }
}

func draw(w, h int) {
    defer termbox.Flush()

    //top bar
    tbPrint(0,0,termbox.ColorBlue, termbox.ColorBlack, "╔")
    for y := 0; y < 1; y++ {
        for x := 1; x < w-1; x++ {
            tbPrint(x, y, termbox.ColorBlue, termbox.ColorBlack, "═")
        }
    }
    tbPrint(w-1,0,termbox.ColorBlue, termbox.ColorBlack, "╗")

    //body

    //bottom bar
    tbPrint(0,h-5,termbox.ColorBlue, termbox.ColorBlack, "╚")
    for y := h-5; y < h-4; y++ {
        for x := 1; x < w-1; x++ {
            tbPrint(x, y, termbox.ColorBlue, termbox.ColorBlack, "═")
        }
    }
    tbPrint(w-1,h-5,termbox.ColorBlue, termbox.ColorBlack, "╝")

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
                        if pX < w-1 && pY < h-1 {
                            pX++
                            pY++
                        } else {
                            pX--
                            pY--
                        }
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
