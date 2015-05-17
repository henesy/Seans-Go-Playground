package main

import (
    "fmt"
    "net"
    "time"
)

type state int
const (
    RUN state = iota
    STOP
)
var toRun state = RUN


func accepter(ln net.Listener, runChan chan state) {
    for {
        conn, err := ln.Accept()
        check(err)
        go handleConnection(conn, runChan)
    }
}

/* runTime reports the time since startTime began */
func runTime(startTime time.Time) {
    endTime := time.Since(startTime)
    fmt.Print("Ran for: ", endTime, "\n")
}

/* handleConnection reads from data separately */
func handleConnection(conn net.Conn, runChan chan state) {
    srvQuit := string([]byte("!srvquit"))

    addr := conn.RemoteAddr()
    fmt.Printf("'%v' connected.\n", addr)
    for {
        srvIn := make([]byte, 512)
        n, err := conn.Read(srvIn)
        if err != nil {
            fmt.Printf("Connection '%v' suffered: '%v'", addr, err)
            break
        }
        srvInString := string(srvIn)
        fmt.Printf("'%d' bytes; Data: '%s'; From: '%v'\n", n, srvInString, addr)
        conn.Write([]byte("Received!\n"))
        if srvInString[:len(srvQuit)] == srvQuit {
            conn.Write([]byte("Closing connection!\n"))
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

/* Simple task manager written in Go, using concurrency; possibly networking */

func main() {
    runChan := make(chan state, 1)

    /* start a master timer to track how long taskmanager ran */
    startTime := time.Now()
    defer runTime(startTime)

    /* begin the listener on port 9090 */
    ln, err := net.Listen("tcp", ":9090")
    check(err)

    go accepter(ln, runChan)

    for running := true; running == true; {
        select {
        case <- runChan:
            close(runChan)
            fmt.Println("STOPPING")
            running = false
        default:
        }
    }
}
