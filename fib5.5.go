package main

import (
	"flag"
	"fmt"
	"os"
	"math/big"
)

var count int = 1

func printer(num1 int, num2 big.Int, countchan chan int, printchan chan int) {
	fmt.Printf("\n%3.0d: %s\n", num1, num2.String())
	printchan <- num1
}

func fib_classic(fibchan chan big.Int, countchan chan int, fibsize int) {
	fib1 := big.NewInt(0)
	fib2 := big.NewInt(1)
	countchan <- 1
	fibchan <- *fib1

	for i := 1; i < fibsize; i += 1 {
		temp2 := new(big.Int)
		temp2.Add(fib1, fib2)
		fib1 = fib2
		fib2 = temp2
		countchan <- i + 1
		fibchan <- *fib1
	}
}

/* efficient fibonacci with "infinite" integer values with concurrency
 modified from fib5 to print only the final number for speed, convenience, and great justice */

func main() {
	var amount int
	flag.IntVar(&amount, "n", 10, "Specify an integer amount of fibonaccis to crunch [2-1477]")
	flag.Parse()
	if amount < 2 {
		fmt.Println("Can only crunch between `> 2` and `< 1477` values.")
		os.Exit(1)
	}
	fibchan := make(chan big.Int, 2)
	countchan := make(chan int, 2)
//	printchan := make(chan int, 2)
	go fib_classic(fibchan, countchan, amount)
	for stahp := false; !stahp; {
		select {
		case num, _ := <-fibchan:
			cnt, _ := <-countchan
//			go printer(cnt, num, countchan, printchan)
//			printcnt := <-printchan
			if cnt == amount {
				fmt.Print(amount, ": ", num.String())
				close(fibchan)
				close(countchan)
//				close(printchan)
				stahp = true
			}
		default:
		}
	}
}
