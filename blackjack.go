package main

import (
"fmt"
"svi"
sc "strconv"
"os/exec"
"os"
)

/* simple blackjack game by sean hinchee */

/* implement splits, bets? dynamic aces?
max cards on board:
max hits: 11 (4xA, 4x2, 3x3) in a hand
max splits: 4 total hands (3 splits)
*/

var deck1 = make([]int, 53)
//var deck1 [52]int
var deck2 = make([]int, 53)
var dboard = make([]string, 53)
var pboard = make([]string, 53)

func initialize_deck_new() {
    num := 2
    for a:=0;a<51;a+=1 {
        deck1[a]=num
        if num >= 11 {
            num=2
        } else {
            num+=1
        }
    }
    for i:=40;i<51;i+=1 {
        deck1[i] = 10
    }
}

func initialize_deck_old() {
    /* initializes deck1 with basic ordered values */

    card_nums := make([]int, 12)
    for h:=1;h<=11;h+=1 { /* 4 of every kind */
        card_nums[h] = 4
    }

    inc := 4.0
    i := 0
    check := 0
    for j:=0;j<51;j+=1 {
        if i == 13 {
            inc+=1
            i=0
        }

        num := svi.Random(1,12)
        for t:=false;t==false; {
            if check >= 52 {break}
            if card_nums[num] < 0 {
                /* num is an available number and can be used for deck */
                t=true
            } else {
                if num > 1 {
                    num=num-1
                } else {
                    if num < 11 {
                        num=num+1
                    } else {
                        num=num-1 /* shouldn't be necessary...but still */
                    }
                }
            }
            check+=1
        }

        deck1[(int(i)+int(inc))] = num
        i+=1
    }
}

func shuffle() {
    /* rand location -> shift rand -> if fail: increment by 1's down then up */
    for i:=0;i<51;i+=1 {
        deck2[i] = deck1[i]
    }

    pos := svi.Random(0,52)
    pos_2 := svi.Random(0,52)
    adj := 4

    for j:=0;j<51;j+=1 {

        num := pos_2 + adj
        if num > 51 {
            num = (num-52)
        }
        //fmt.Print("DEBUG: ", num, pos)
        deck2[pos] = deck1[num]

        if pos > 51 {
            pos=(pos-51)
        } else {
            pos+=1
        }

            pos_2+=1
    }

}

func deal() {
    i:=0
    for ;i<2;i+=1 {
        pboard[i] = sc.Itoa(deck2[i])
    }
    i=0
    for j:=3;j<5;j+=1 {
        dboard[i] = sc.Itoa(deck2[j])
        i+=1
    }
}

func board(num_card int, cards ...string)() {
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
   cmd := exec.Command("clear")
   cmd.Stdout = os.Stdout
   //c := "╔══╗\n║%s║\n╚══╝\n"

   fmt.Print("Would you like instructions? [y/n] ")
   fmt.Scanln(&in)

   cmd.Run()
   n_d:=2
   n_p:=2
   for run:=true;run!=false; {
       var usrin string
       fmt.Scanln(&usrin)

       initialize_deck_new()
       fmt.Print(deck1)
       shuffle()
       deal()
       fmt.Print("\nDealer Board\n")
       board(n_d, dboard...)
       fmt.Print("\nPlayer Board\n")
       board(n_p, pboard...)
       fmt.Print("\nCommand: ")

       if usrin == "quit" {
           run=false
       }
   }
   fmt.Println("Good Bye!")
}
