package main

import (
"fmt"
)

/* simple blackjack game by sean hinchee */

/* implement splits, bets? 
max cards on board: 
max hits: 11 (4xA, 4x2, 3x3) in a hand
max splits: 4 total hands (3 splits)
*/

func card(n, i string)(string) {
    var s string
    switch i {
        case "s": s = "♠"
        case "h": s = "♥"
        case "d": s = "♦"
        case "c": s = "♣"
        default: s = "🃏"
    }
    
    /* output ascii image 
    ╔══╗
    ║q♥║
    ╚══╝ */

    return n+s
}

func main() {
   var in int
   c := "╔══╗\n║%s║\n╚══╝\n"
   //dboard := make([]string, 52)
   //pboard := make([]string, 52)
   
   fmt.Print("Would you like instructions? [y/n] ")
   fmt.Scanln(&in)
   
   l := card("q", "s")
   fmt.Printf(c, l)

}

