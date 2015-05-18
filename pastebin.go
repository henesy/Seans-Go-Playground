/* copied as a useful reference from a pastebin link i ran across
with termbox-go examples */
package main

import (
        "fmt"
        "github.com/nsf/termbox-go"
)

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
        for _, c := range msg {
                termbox.SetCell(x, y, c, fg, bg)
                x++
        }
}

func draw(playerX int, playerY int) {
        termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
        defer termbox.Flush()

        w, h := termbox.Size()

        startX := playerX - w / 2
        startY := playerY - h / 2

        for x := 0; x < w; x++ {
                for y := 0; y < h; y++ {
                        tbPrint(x, y, termbox.ColorBlue, termbox.ColorDefault, "~")
                }
        }

        // The player
        tbPrint(w/2, h/2, termbox.ColorRed, termbox.ColorDefault, "O")

        // The static object
        tbPrint(90 - startX, 90 - startY, termbox.ColorRed, termbox.ColorDefault, "X")
}

func main() {
        err := termbox.Init()
        if err != nil {
                panic(err)
        }
        termbox.SetInputMode(termbox.InputEsc)

        playerX := 100
        playerY := 100
        draw(playerX, playerY)

        mainLoop:
        for {
                switch ev := termbox.PollEvent(); ev.Type {
                        case termbox.EventKey:
                                key := string(ev.Ch)

                                if key == "q" {
                                        break mainLoop
                                }

                                if key == "w" {
                                        playerY--
                                }

                                if key == "s" {
                                        playerY++
                                }

                                if key == "a" {
                                        playerX--
                                }

                                if key == "d" {
                                        playerX++
                                }

                                draw(playerX, playerY)
                }
        }

        termbox.Close()

        fmt.Println("Done")
}
