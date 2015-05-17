package main

import (
    "fmt"
    "net"
    "time"
)

/* runTime reports the time since startTime began */
func runTime(startTime time.Time) {
    endTime := time.Since(startTime)
    fmt.Print("Ran for: ", endTime, "\n")
}

/* handleConnection reads from data separately */
func handleConnection(conn net.Conn) {
    srvIn := make([]byte, 512)
    addr := conn.RemoteAddr()
    fmt.Printf("'%v' connected.\n", addr)
    for {
        n, err := conn.Read(srvIn)
        if err != nil {
            fmt.Printf("Connection '%v' suffered: '%v'", addr, err)
            break
        }
        fmt.Printf("'%d' bytes; Data: '%s'; From: '%v'\n", n, string(srvIn), addr)
        conn.Write([]byte("Received!\n"))
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
    /* start a master timer to track how long taskmanager ran */
    startTime := time.Now()
    defer runTime(startTime)

    /* begin the listener on port 9090 */
    ln, err := net.Listen("tcp", ":9090")
    check(err)

    for {
        conn, err := ln.Accept()
        check(err)
        go handleConnection(conn)
    }
}
