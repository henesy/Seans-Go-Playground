package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Int() % 8191)

	for {
		in := ""
		fmt.Scanln(&in)
		fmt.Println("?")
	}
}
