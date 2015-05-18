package main

import (
    "fmt"
    "net"
    "time"
    "flag"
)

//go:generate go build -o shclient client/shclient.go
//go:generate go build -o socketh

type state int
const (
    RUN state = iota
    STOP
)
type connected int
const (
    DIS connected = iota
    CON
)
var toRun state = RUN
var max uint64 = 100
var connections []net.Conn = make([]net.Conn, max)
var cState []connected = make([]connected, max)
var usrNames []string = make([]string, max)
var numConns, maxConns uint64
var srvPos uint64 = max+1

/* messages all connected clients when messages are sent */
func messageConns(pos uint64, addr, words string) {
    var i uint64
    for i = 0; i < maxConns;i += 1 {
        if cState[i] == CON && i != pos {
            message := "\n" + addr + ": " + words
            connections[i].Write([]byte(message))
        }
    }
}

/* finds open connection position */
func findOpen()(pos uint64) {
    var i uint64
    for i = 0;i < maxConns;i += 1 {
        if cState[i] == DIS {
            pos = i
        }
    }
    return
}

/* close connections */
func closeConnection(pos uint64) {
    cState[pos] = DIS
}

/* acceps connections, kept spinning separate from main() */
func accepter(ln net.Listener, runChan chan state) {
    for {
        if numConns < maxConns {
            conn, err := ln.Accept()
            check(err)
            pos := findOpen()
            connections[pos] = conn
            cState[pos] = CON
            go handleConnection(&connections[pos], runChan, time.Now(), pos)
            numConns++
        } else {
            ln.Close()
            fmt.Print("Max connections reached.\n")
        }
    }
}

/* runTime reports the time since startTime began */
func runTime(name string, startTime time.Time) {
    endTime := time.Since(startTime)
    fmt.Printf("%s was connected for: %v\n", name, endTime)
}

/* handleConnection reads from data separately */
func handleConnection(conn *net.Conn, runChan chan state, rTime time.Time, pos uint64) {
    srvQuit := string([]byte("!srvquit"))
    exQuit := string([]byte("!quit"))
    defer closeConnection(pos)

    tmpUsrName := make([]byte, 25)
    strUsrName := ""
    (*conn).Write([]byte("What is your username?: "))
    (*conn).Read(tmpUsrName)
    cnt := 0
    for i := 0;i < len(tmpUsrName); i += 1 {
        if tmpUsrName[i] == byte(0) {
            cnt++
        }
        if cnt > 3 {
            strUsrName = (string(tmpUsrName[:i-cnt+1]))
        }
    }
    usrNames[pos] = strUsrName

    addr := (*conn).RemoteAddr()
    fmt.Printf("'%v' connected.\n", addr)
    go messageConns(srvPos, usrNames[pos], "→ Connected.")

    for {
        srvIn := make([]byte, 512)
        n, err := (*conn).Read(srvIn)
        if err != nil {
            go messageConns(srvPos, usrNames[pos], "← Disconnected.")
            fmt.Printf("Connection '%v' suffered: '%v'\n", addr, err)
            runTime((*conn).RemoteAddr().String(), rTime)
            numConns--
            break
        }
        srvInString := string(srvIn)
        fmt.Printf("'%d' bytes; Data: '%s'; From: '%v'\n", n, srvInString, addr)
        if srvInString[:len(exQuit)] == exQuit {
            go messageConns(srvPos, usrNames[pos], "← Disconnected.")
            runTime((*conn).RemoteAddr().String(), rTime)
            numConns--
            break
        }
        go messageConns(pos, usrNames[pos], srvInString)
        (*conn).Write([]byte("Received!"))
        if srvInString[:len(srvQuit)] == srvQuit {
            go messageConns(srvPos, (*conn).LocalAddr().String(),"Closing connection!")
            runTime((*conn).RemoteAddr().String(), rTime)
            runChan <- STOP
        }
    }
}

/* check checks the error err for an error and crashes the program if != nil */
func check(err error) {
    if err != nil {
        panic(err)
    }
}

/* Simple chat server written in Go, using concurrency */

func main() {
    var port string
    flag.Uint64Var(&maxConns, "mc", 30, "Set the max number of connections.")
    flag.StringVar(&port, "p", ":9090", "Set the port to listen on.")
    flag.Parse()
    runChan := make(chan state, 1)
    //messageChan := make(chan string, 3)

    /* start a master timer to track how long taskmanager ran */
    startTime := time.Now()
    defer runTime("main", startTime)

    /* begin the listener on port 9090 */
    ln, err := net.Listen("tcp", port)
    check(err)

    go accepter(ln, runChan)

    for running := RUN; running == RUN; {
        select {
        case <- runChan:
            close(runChan)
            fmt.Println("STOPPING")
            running = STOP
        default:
        }
        time.Sleep(1 * time.Second)
    }
}
