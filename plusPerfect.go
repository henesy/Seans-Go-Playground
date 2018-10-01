package main

import (
	"fmt"
	"math"
	"os"
sc	"strconv"
)

func isPlusPerfect(n float64) bool {
	mag := int(math.Floor(math.Log10(n)))
	var pp float64 = 0

	for i := mag; i >= 0; i-- {
		pp += math.Pow(float64(int(n / math.Pow10(i)) % 10), float64(mag+1))
	}

	return pp == n
}

/* Takes a list of integers as arguments and tests if they are plus perfect or not */
func main() {
	usage := func() {
		fmt.Println("usage: ", os.Args[0], " nums...")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		usage()
	}

	for _, str := range os.Args[1:] {
		// Trust
		nint, _ := sc.Atoi(str)
		n := float64(nint)
		fmt.Print(n)
		if isPlusPerfect(n) {
			fmt.Println(" is plus perfect")
		} else {
			fmt.Println(" is not plus perfect")
		}
	}
}

