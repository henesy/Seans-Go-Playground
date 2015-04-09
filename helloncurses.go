package main

import (
    "github.com/rthornton128/goncurses"
    "log"
)

/* using https://github.com/rthornton128/goncurses */

func main() {
    stdscr, err := goncurses.Init()
    if err != nil {
        log.Fatal("init", err)
    }
    defer goncurses.End()

    stdscr.Print("Hello, World!")
    stdscr.MovePrint(3, 0, "Press any key to continue")

    stdscr.Refresh()

    stdscr.GetChar()
}

