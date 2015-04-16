package main

import (
"fmt"
"svi"
sc "strconv"
)

/* simple blackjack game by sean hinchee */

/* implement splits, bets? dynamic aces?
max cards on board:
max hits: 11 (4xA, 4x2, 3x3) in a hand
max splits: 4 total hands (3 splits)
*/

var deck = make([]string, 52)
//dboard := make([]string, 50)
var pboard = make([]string, 50)

func shuffle() {
    /* this seriously needs a proper algorithm
    and/or array-ified (pointers?)
    perhaps shuffle backwards and decrement in a big array?*/
    var p float64
    spades := 0
    diamonds := 0
    hearts := 0
    clubs := 0
    tens := 0
    jacks := 0
    queens := 0
    kings := 0
    aces := 0

    nums := make([]int, 12)
    for i:=0; i<12;i+=1{
        nums[i] = 0
    }

    z:=0
    for shuffled:=false;shuffled==false; {

        a := svi.Random(1, 11) //Numbered card selection
        b := svi.Random(0, 4) //Face card selection; T,j,q,k,A
        c := svi.Random(0, 4) //Card suit selection; s,d,h,c
        d := "A" //Left out
        e := "*" //Right out

        //for y:=0;y<1; {
            if a == 11 && aces<4 {
                d = "A"
                aces+=1
                //y+=1
            } else if a == 10 {
                if b == 0 && tens<4 {
                    d = "T"
                    tens+=1
                    //y+=1
                } else if b == 1 && jacks<4 {
                    d = "J"
                    jacks+=1
                    //y+=1
                } else if b == 2 && queens<4 {
                    d = "Q"
                    queens+=1
                    //y+=1
                } else if b == 3 && kings<4 {
                    d = "K"
                    kings+=1
                    //y+=1
                } else {
                    continue
                }
            } else {
                if (nums[a] / a) < 4 {
                    d = sc.Itoa(a)
                    nums[a] = nums[a] + a
                    //y+=1
                } else {
                    continue
                }
            }

            if c == 0 && spades<13 {
                e = "s"
                spades+=1
                //y+=1
            } else if c == 1 && diamonds<13 {
                e = "d"
                diamonds+=1
                //y+=1
            } else if c == 2 && hearts<13 {
                e = "h"
                hearts+=1
                //y+=1
            } else if c == 3 && clubs<13 {
                e = "c"
                clubs+=1
                //y+=1
            } else {
                continue
            }
        //}


        if z == 52 {
            break
        }

        deck[z] = card(d, e)
        z+=1.0 //does not increment if continue is hit
        p = ((float64(z)+48)/(52+48))*100
        go fmt.Printf("\r%0.0f%% Shuffled", p)

        /* choose face/num -> */
        //ncard = svi.Random(1, 10)
    }

}

func board(num_card int, cards ...string)() {
    fmt.Print("\nPlayer Board\n")

    for i:=0; i<80; i+=1 {
        fmt.Printf("―")
    }

    for i:=0; i<num_card; i+=1 {
        fmt.Printf("╔══╗")
    }

    fmt.Printf("\n")

    z:=0
    for _, face := range cards {
        if z == num_card {
            break
        }
        fmt.Printf("║%s║", face)
        z+=1
    }

    fmt.Printf("\n")

    for i:=0; i<num_card; i+=1 {
        fmt.Printf("╚══╝")
    }

    fmt.Print("\nCommand: ")
}

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
   //c := "╔══╗\n║%s║\n╚══╝\n"

   fmt.Print("Would you like instructions? [y/n] ")
   fmt.Scanln(&in)

   n:=2
   for usrin:="" ; usrin != "quit"; {

       shuffle()
       pboard[0] = deck[0]
       pboard[1] = deck[1]
       go board(n, pboard...)
       fmt.Scanln(&usrin)

       //l := card("q", "s")


   }
}
