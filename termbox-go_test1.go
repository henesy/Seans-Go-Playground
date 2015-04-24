package main

import(
    "fmt"
    "github.com/nsf/termbox-go"
)

/* termbox-go test for fun and profit */

const grass = termbox.ColorGreen
const coldef = termbox.ColorDefault
//var smooth rune = '#'


func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func main() {
    //data := make([]byte, 0, 64)

    err := termbox.Init()
    if err != nil {
        fmt.Print(err, "\n")
    }
    defer termbox.Close()
    termbox.SetInputMode(termbox.InputAlt)
    termbox.Clear(coldef, coldef)
    termbox.Flush()
    //termbox.SetCell(0, 0, smooth, termbox.ColorGreen, termbox.ColorBlack)
    fill(0,0,80,24,termbox.Cell{Ch: '▓'})
    termbox.Flush()
    tbprint(12, 40, termbox.ColorCyan, grass, "Welcome to [Game Name Here]")
    termbox.Flush()
    for ;; {
        switch event := termbox.PollEvent(); event.Type {
            case termbox.EventKey:
                tbprint(12, 40, termbox.ColorBlue, termbox.ColorMagenta, string(event.Ch))
                if event.Ch == 'q' {
                    break
                }
            default:
                tbprint(12, 40 , termbox.ColorRed, termbox.ColorBlack, "No...no...")
        }
        termbox.Flush()
    }

}
