package main

import (
"fmt"
//"math/cmplx"
)

type MSegment struct {
p1 string //= "•" // dot
p2 string //= "-" // line
p3r string //= "→" // right finalizer
p3l string //= "←" // left finalizer
}

type MRow struct {
	v1 float64 // start velocity
	v2 float64 // final velocity
	x1 float64 // start point
	x2 float64 // endpoint
	t float64 // time increment per dot
}

var (
x1 complex128 // previous position; original first pos
x2 complex128 // current pos
x3 complex128 // next pos
y float64 // generic placeholder
xposz int // position changes
)

func (s MSegment) (numOf int, direction string) {
	if direction == "r" {
		for n:=0; n < numOf; n+=1 {
			fmt.Println(s.p1)
			fmt.Println("._.")
		}
	} else {
		fmt.Println("meep")
	}
	fmt.Println("test")
}



func main() {
	seg := MSegment{"•", "-", "→", "←"}

	fmt.Printf("Number of position changes?: ")
	fmt.Scan(&xposz)
	fmt.Printf("Using %v for number of position changes.\n", xposz)
	
	fmt.Printf("Time per dot (no label): ")
	fmt.Scan(&y)

	fmt.Printf("Starting displacement: ")
	fmt.Scan(&x1)
	x2 = x1
	
	fmt.Printf("Ending displacement: ")
	fmt.Scan(&x3)

	fmt.Printf("This: \n%v", x1, x2, x3, y, xposz)
	
}