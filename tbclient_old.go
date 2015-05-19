package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
    "flag"
    "unicode/utf8"
    "net"
    "io"
    "time"
sc  "strconv"
)

/* message buffer to print to screen and possibly scroll in the future */
var msgs []string = make([]string, 50)
var posMsgs uint32 = 0
/* text typed by the user */
var txt []byte = make([]byte, 1024)
var posTxt uint32 = 0

var pos, lpos int

/* populates messages with blank messages */
func popMsgs() {
    for pos := range msgs {
        msgs[pos] = " "
    }
}

/* clears user buffer */
func clearTxt() {
    for pos := range txt {
        txt[pos] = byte(' ')
    }
}

/* prints to pos x, y */
func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
    for _, c := range msg {
            termbox.SetCell(x, y, c, fg, bg)
            x++
    }
}

/* draws the screen */
func draw(w, h int) {
    for {
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
        i := 0+posMsgs
        for y := 1; y < h - 5; y++ {
            tbPrint(0, y, termbox.ColorBlue, termbox.ColorBlack, "║")
            tbPrint(w-1, y, termbox.ColorBlue, termbox.ColorBlack, "║")
        }

        /* this code is so shit */
        for y := 1; y < h - 5; y++ {
            msgs[0] = "DecodeRuneInString is like DecodeRune but its input is a string. If s is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it returns (RuneError, 1). Both are impossible results for correct "
            tot0 := utf8.RuneCountInString(msgs[0])
            msgs[3] = sc.Itoa(tot0)

            tot := utf8.RuneCountInString(msgs[i])
            if tot < 78 {
                msgs[1] = "sup"
                msgs[2] = "test"
                tbPrint(1, y, termbox.ColorCyan, termbox.ColorBlack, msgs[i])
                i += 1
            /* not running atm, handles msgs > 78 chars long */
            } else {
                tmpmsg := msgs[i]
                tmpch := make([]rune, tot)
                for h := 0; h < tot ; h++ {
                    char, size := utf8.DecodeRuneInString(tmpmsg)
                    tmpch[h] = char
                    tmpmsg = tmpmsg[size:]
                }
                msgs[2] = string(tmpch)

                i += 1
            }
        }
        /* end bad code */

        //bottom bar
        tbPrint(0,h-5,termbox.ColorBlue, termbox.ColorBlack, "╚")
        for y := h-5; y < h-4; y++ {
            for x := 1; x < w-1; x++ {
                tbPrint(x, y, termbox.ColorBlue, termbox.ColorBlack, "═")
            }
        }
        tbPrint(w-1, h-5, termbox.ColorBlue, termbox.ColorBlack, "╝")

        //user input zone !-! should make this a for loop
        pos=w-1
        lpos=0
        tbPrint(0, h-4, termbox.ColorWhite, termbox.ColorBlack, string(txt[lpos:pos]))
        lpos=pos
        pos+=w-1
        tbPrint(0, h-3, termbox.ColorWhite, termbox.ColorBlack, string(txt[lpos:pos]))
        lpos=pos
        pos+=w-1
        tbPrint(0, h-2, termbox.ColorWhite, termbox.ColorBlack, string(txt[lpos:pos]))
        lpos=pos
        pos+=w-1
        tbPrint(0, h-1, termbox.ColorWhite, termbox.ColorBlack, string(txt[lpos:pos]))
        termbox.Flush()
        time.Sleep(20 * time.Millisecond)
    }
}

/* reads things from the server and handles them */
func readServer(conn net.Conn, runChan chan bool, pChan chan uint32) {
    //var ticker uint32 = 0
    for {
        words := make([]byte, 512)
        _, err := conn.Read(words)
        if err == io.EOF {
            fmt.Print("\nDisconnected from server.\n")
            runChan <- false
            break
        } else {
            check(err)
        }

        blank := []byte{0,0,0,0,0,0}
        cnt := 0
        for p, w := range blank {
            if words[p] == w {
                cnt++
            }
        }
        if cnt < 4 {
            fmt.Print(string(words), "\n")
            pChan <- uint32(1)
            time.Sleep(20 * time.Millisecond)
        }
    }
}

/* dialServer will connect to a pre-selected server */
func dialServer(target string) {
    runChan := make(chan bool, 1)
    pChan := make(chan uint32, 1)
    //var words []byte
    conn, err := net.Dial("tcp", target)
    if err != nil {
        fmt.Print(err, "\n")
    }
    /* get our info back from the server */
    go readServer(conn, runChan, pChan)
    //go readIn(conn, runChan, pChan)

    for run := true; run == true; {
        select {
        case <- runChan:
            close(runChan)
            close(pChan)
            run = false
        default:
        }
        time.Sleep(20 * time.Millisecond)
    }
}

/* check checks the error err for an error and crashes the program if != nil */
func check(err error) {
    if err != nil {
        panic(err)
    }
}

/* client for SocketH rewritten in termbox-go */

func main() {
    words := ""
    flag.StringVar(&words, "a", "localhost:9090", "Set address to dial to for server.")
    flag.Parse()
    defer fmt.Print("\nGoodbye!\n")
    defer termbox.Close()
    popMsgs()
    //go dialServer(words)


    err := termbox.Init()
    if err != nil {
        fmt.Println(err)
    }
    termbox.SetInputMode(termbox.InputAlt)
    w, h := termbox.Size()

    termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
    termbox.Flush()
    go draw(w,h)
    termbox.Flush()
    for run:=true;run==true; {
        switch ev := termbox.PollEvent(); ev.Type {
            case termbox.EventKey:
                key := string(ev.Ch)

                    if ev.Key == termbox.KeyCtrlQ {
                        run=false
                        termbox.Flush()
                    } else if ev.Key == termbox.KeyEnter {
                        posMsgs++
                        oldtxt := txt
                        clearTxt()

                        posTxt = 0
                        tmptxt := ""
                        tick := 0
                        for len(oldtxt) > 0 {
		                  r, size := utf8.DecodeRune(oldtxt)
                          if r == ' ' {
                              tick+=1
                          }
                          if tick > 3 {
                              break
                          }
                          tmptxt += string(r)
		                  oldtxt = oldtxt[size:]
	                   }
                       msgs[posMsgs] = " "
                       msgs[posMsgs] = "wat"

                    } else if ev.Key == termbox.KeyBackspace2 || ev.Key == termbox.KeyBackspace {
                        txt[posTxt] = byte(' ')
                        txt[posTxt-1] = byte(' ')
                        posTxt-=1
                    } else {
                        letr, _ := utf8.DecodeRuneInString(key)
                        txt[posTxt] = byte(letr)
                        posTxt++
                    }
        }
    }
}
