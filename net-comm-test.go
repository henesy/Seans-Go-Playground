package main

import (
	"fmt"
	"net"
	"encoding/json"
)

/* a person?! */
type Person struct {
	Age		int
	Name	string
}

/* check and print errors */
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/* tests a basic end-to-end network tranmission of a structure */
func main() {
	/* run one instance as server and one as client to simulate connection */
	P:
	sorc := ""
	fmt.Print("[c]lient or [s]erver?: ")
	fmt.Scanln(&sorc)
	switch sorc {
		case "c":
			conn, err := net.Dial("tcp", "localhost:5573")
			check(err)
			enc := json.NewEncoder(conn)
			Mvp := Person{18, "michael!"}
			err = enc.Encode(Mvp)
			check(err)
		case "s":
			ln, err := net.Listen("tcp", ":5573")
			check(err)
			conn, err := ln.Accept()
			check(err)
			dec := json.NewDecoder(conn)
			var Michael Person
			err = dec.Decode(&Michael)
			check(err)
			fmt.Println(Michael)
		case "q":
			break
		default: goto P
	}

}

