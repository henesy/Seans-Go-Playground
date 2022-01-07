package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	nPrimes := 100
	primes := map[int]bool{
		2: true,
	}
	fmt.Fprintln(w, 2)

	for i := 2; ; i++ {
		if isPrime(primes, i) {
			primes[i] = true
			fmt.Fprintln(w, i)
		}
		if len(primes) >= nPrimes {
			break
		}
	}
}

func isPrime(primes map[int]bool, x int) bool {
	for p := range primes {
		if x%p == 0 {
			return false
		}
	}
	return true
}
