package main

import(
"fmt"
"net"
)

/* simple proof of concept for go's "net" package */

func handleConnection(cn Conn) {
    var
    cn.Read

}

func main() {
   ln, err := net.Listen("tcp", ":9999")
   if err != nil {
        // handle error
   }
   for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
        }
        go handleConnection(conn)
   }


}
