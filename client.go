package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
)

func readServer(conn net.Conn) {
    words := make([]byte, 512)
    for {
        _, err := conn.Read(words)
        check(err)
        fmt.Print(string(words))
    }
}

/* dialServer will connect to a pre-selected server */
func dialServer(target string) {
    //var words []byte
    conn, err := net.Dial("tcp", target)
    if err != nil {
        fmt.Print(err, "\n")
    }
    /* get our info back from the server */
    go readServer(conn)

    reader := bufio.NewReader(os.Stdin)
    for {
        words, _, _ := reader.ReadLine()
        _, err := conn.Write(words)
        check(err)
        if string(words) == "!quit" {
            break
        }
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
