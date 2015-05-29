package main

import (
	"fmt"
	"time"
)

type lb uint32
var pos int
const LOAD lb = iota

/* loadinz */
func loader(loadChan chan lb, bar []rune) {
	for {
		bar[pos] = '#'
		loadChan <- 0
		time.Sleep(time.Duration(1 * time.Second))
	}
}

/* a simple demonstration program on loading bars */

func main() {
	loadChan := make(chan lb, 1)
	bar := make([]rune, 78)

	for h := 0; h < len(bar); h++ {
		bar[h] = ' '
	}

	go loader(loadChan, bar)

	for pos = 0; pos < len(bar); pos++ {
		select {
			case <- loadChan:
				
				fmt.Print("\r[")
				for j := 0; j < len(bar); j++ {
					fmt.Printf("%c", bar[j])
				}
				fmt.Print("]")

		}
	}
	fmt.Print("\n")
	close(loadChan)
}
