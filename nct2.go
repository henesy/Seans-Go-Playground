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
			Mvp1 := Person{18, "michael!"}
			Mvp2 := Person{20, "brandon~"}
			Mvp3 := Person{99, "kenpachiramasama"}
			Mvps := make([]Person, 3)
			Mvps[0], Mvps[1], Mvps[2] = Mvp1, Mvp2, Mvp3
			err = enc.Encode(Mvps)
			check(err)
		case "s":
			ln, err := net.Listen("tcp", ":5573")
			check(err)
			conn, err := ln.Accept()
			check(err)
			dec := json.NewDecoder(conn)
			Peeps := make([]Person, 3)
			for dec.More() {
				err = dec.Decode(&Peeps)
				check(err)
			}
			fmt.Println(Peeps)
		case "q":
			break
		default: goto P
	}

}

