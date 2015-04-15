package main

import(
"fmt"
"github.com/rthornton128/goncurses"
)

/* simple blackjack game by sean hinchee using goncurses */

/* implement splits, bets? 
max cards on board: 
max hits: 11 (4xA, 4x2, 3x3) in a hand
max splits: 4 total hands (3 splits)
*/

func pcard(i n string)(o string) {
    /* icon/suit */
    switch i {
        case "s": s := "♠"
        case "h": s := "♥"
        case "d": s := "♦"
        case "c": s := "♣"
        default: s := "🃏"
    }
    
    /* output ascii image 
    ╔══╗
    ║q♥║
    ╚══╝
    */
    o := ""
    return
}

func main() {
   var in int
   dboard = make([]string, 52)
   pboard = make([]string, 52)
   stdscr, err := goncurses.Init()
   defer goncurses.End()
   
   stdscr.Print("Would you like instructions? [y/n] ")
   stdscr.Refresh()
   stdscr.Getchar(&in)


}

