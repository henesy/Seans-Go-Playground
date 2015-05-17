package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "io"
    "time"
)

func readIn(conn net.Conn, runChan chan bool) {
    reader := bufio.NewReader(os.Stdin)
    for {
        words, _, _ := reader.ReadLine()
        _, err := conn.Write(words)
        check(err)
        if string(words) == "!quit" {
            runChan <- false
            break
        }
    }
}

func readServer(conn net.Conn, runChan chan bool) {
    for {
        words := make([]byte, 512)
        _, err := conn.Read(words)
        if err == io.EOF {
            fmt.Print("Disconnected from server.\n")
            runChan <- false
            break
        } else {
            check(err)
        }
        fmt.Print(string(words), "\n")
    }
}

/* dialServer will connect to a pre-selected server */
func dialServer(target string) {
    runChan := make(chan bool, 1)
    //var words []byte
    conn, err := net.Dial("tcp", target)
    if err != nil {
        fmt.Print(err, "\n")
    }
    /* get our info back from the server */
    go readServer(conn, runChan)
    go readIn(conn, runChan)

    for run := true; run == true; {
        select {
        case <- runChan:
            run = false
        default:
        }
        time.Sleep(1 * time.Second)
    }
}

/* check checks the error err for an error and crashes the program if != nil */
func check(err error) {
    if err != nil {
        panic(err)
    }
}

/* Simple raw connection client */

func main() {
    defer fmt.Print("Goodbye!\n")
    fmt.Print("Dial address?: ")
    words := ""
    fmt.Scanln(&words)
    fmt.Print("Dialing ", words, ";\n")
    dialServer(words)
}
