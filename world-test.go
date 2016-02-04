package main

import (
	"fmt"
	"net"
	"encoding/json"
sc	"strconv"
	"strings"
	"os"
)

type TicTacToe [9]Sprite
type Sprites []Sprite
/* sprite struct for world indexing */
type Sprite struct {
	R	rune
	X	int
	Y	int
}

/* error checker and printer */
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/* create and return a world */
func genWorld(size int)(Sprites) {
	World := make([]Sprite, size)

	//makes tic tac toe
	y := 0
	for p, _ := range World {
		World[p].X = p % 3
		World[p].Y = y
		if World[p].X == 2 {
			y++
		}
		World[p].R = ' '
	}
	return Sprites(World)
}

/* places an X at a designed coordinate */
func (W* Sprites) X(x, y int) {
	for p, v := range (*W) {
		if (v).X == x && (v).Y == y {
			(*W)[p].R = 'X'
		}
	}
}

/* places an O at a designed coordinate */
func (W* Sprites) O(x, y int) {
	for p, v := range (*W) {
		if (v).X == x && (v).Y == y {
			(*W)[p].R = 'O'
		}
	}
}

/* prompt, calls encoder */
func prompt(conn net.Conn, World Sprites, t string, s int) {
	rply := ""
	fmt.Print("X,Y: ")
	fmt.Scanln(&rply)
	if rply == "q" {
		os.Exit(1)
	}
	co := strings.Split(rply, ",")
	xs, ys := co[0], co[1]
	x, _ := sc.Atoi(xs)
	y, _ := sc.Atoi(ys)

	if t == "s" {
		(&World).X(x, y)
	} else {
		(&World).O(x, y)
	}
	encoder(conn, World, t, s)
}

/* decodes, calls prompt */
func decoder(conn net.Conn, World Sprites, t string, s int) {
	World2 := make([]byte, s)
	conn.Read(World2)
	json.Unmarshal(World2, &World)
	curY := 0
	for _, v := range World {
		if v.Y > curY {
			fmt.Print("\n")
			curY = v.Y
		}
		fmt.Printf("%c", v.R)
	}
	fmt.Print("\n")
	prompt(conn, World, t, s)
}

/* encoder, calls decoder*/
func encoder(conn net.Conn, World Sprites, t string, s int) {
	W, _ := json.Marshal(World)
	s, _ = conn.Write(W)
	decoder(conn, World, t, s)
}


/* a test in copying a "world" over the network for processing */
func main() {

	P:
	sorc := ""
	fmt.Print("[s]erver or [c]lient?: ")
	fmt.Scanln(&sorc)
	switch sorc {
		case "c":
			World := genWorld(9)
			//World[0].R = 'X'
			//(&World).O(1, 1)
			conn, err := net.Dial("tcp", "localhost:5573")
			check(err)
			prompt(conn, World, sorc, 190)
		case "s":
			World := genWorld(9)
			ln, err := net.Listen("tcp", ":5573")
			check(err)
			conn, err := ln.Accept()
			check(err)
			decoder(conn, World, sorc, 190)
		case "q": break
		default: goto P
	}

	fmt.Println("Goodbye!")
}




