package main

import (
	"fmt"
	"flag"
)

/* inefficient capped fibonacci */

func main() {
	var N int
	flag.IntVar(&N, "n", 1000, "Perform `N` number of iterations [1000]")
	flag.Parse()
	fibs := make([]float64, N)
	fibs[0]=0.0
	fibs[1]=1.0
	for i:=2;i<N;i++ {
		nums := fibs[i-2:i]
		fibs[i] = nums[0] + nums[1]
	}
	fmt.Printf("%d: %e\n", N, fibs[N-1])
}

