package main

import (
	"flag"
	"fmt"
	"os"
)

var count int = 0

func printer(num1 *int, num2 *float64, countchan chan *int, printchan chan *int) {
	fmt.Printf("%3.0d: %.0f\n", *num1-2, *num2)
	printchan <- num1
}

func fib_roundabout(fibchan chan *float64, countchan chan *int, fibsize *int) {
	fibs := make([]*float64, *fibsize)
	var a, b float64 = 0.0, 1.0
	fibs[0], fibs[1] = &a, &b

	count = 1
	fibchan <- fibs[0]
	countchan <- &count

	count = 2
	fibchan <- fibs[1]
	countchan <- &count

	for i := 2; i < *fibsize; i += 1 {
		count = (i+1)
		nums := fibs[i-2 : i]
		var c, d float64 = *nums[0], *nums[1]
		e := d + c
		fibs[i] = &e
		countchan <- &count
		fibchan <- fibs[i]
	}
}

/* worked before pointers, needs updating/fixing */
func fib_classic(fibchan chan *float64, countchan chan *int, fibsize *int) {
	var fib1, fib2 float64
	var a int = 1
	fib1, fib2 = 0, 1

	countchan <- &a
	fibchan <- &fib1

	for i := 1; i < *fibsize; i += 1 {
		fib1, fib2 = fib2, (fib1 + fib2)
		countchan <- &i
		fibchan <- &fib1
	}
}

/* inefficient capped fibonacci with pointers and concurrency */

func main() {
	var amount int
	flag.IntVar(&amount, "n", 12, "Specify an integer amount of fibonaccis to crunch [2-1477]")
	flag.Parse()
	if amount < 2 {
		fmt.Println("Can only crunch between `> 2` and `< 1477` values. Program adds 2.")
		os.Exit(1)
	}
	amount+=2
	fibchan := make(chan *float64, 1)
	countchan := make(chan *int, 1)
	printchan := make(chan *int, 1)
	go fib_roundabout(fibchan, countchan, &amount)
	for stahp := false; !stahp; {
		select {
		case num, _ := <-fibchan:
			cnt, _ := <-countchan
			go printer(cnt, num, countchan, printchan)
			printcnt := <-printchan
			if *cnt == amount && *printcnt == amount {
				close(fibchan)
				close(countchan)
				close(printchan)
				stahp = true
			}
		default:
		}
	}
}
