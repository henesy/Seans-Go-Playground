package main

import (
	"flag"
	"fmt"
	"os"
)

//fmt.Printf("%3d: %.0f\n", i+1, fibs[i])
var count int = 1

func printer(num1 int, num2 float64, countchan chan int, printchan chan int) {
	fmt.Printf("%3.0d: %.0f\n", num1, num2)
	printchan <- num1
}

func fib_roundabout(fibchan chan float64, countchan chan int, fibsize int) {
	fibs := make([]float64, fibsize)
	fibs[0], fibs[1] = 0, 1
	fibchan <- fibs[0]
	countchan <- count
	fibchan <- fibs[1]
	count += 1
	countchan <- count

	for i := 2; i < fibsize; i += 1 {
		count = i + 1
		nums := fibs[i-2 : i]
		fibs[i] = nums[0] + nums[1]
		countchan <- count
		fibchan <- fibs[i]
	}
}

func fib_classic(fibchan chan float64, countchan chan int, fibsize int) {
	var fib1, fib2 float64
	fib1, fib2 = 0, 1
	countchan <- 1
	fibchan <- fib1
	for i := 1; i < fibsize; i += 1 {
		fib1, fib2 = fib2, fib1+fib2
		countchan <- i + 1
		fibchan <- fib1
	}
}

func main() {
	var amount int
	flag.IntVar(&amount, "n", 10, "Specify an integer amount of fibonaccis to crunch [2-1477]")
	flag.Parse()
	if amount < 2 {
		fmt.Println("Can only crunch between `> 2` and `< 1477` values.")
		os.Exit(1)
	}
	fibchan := make(chan float64, 2)
	countchan := make(chan int, 2)
	printchan := make(chan int, 2)
	go fib_classic(fibchan, countchan, amount)
	for stahp := false; !stahp; {
		select {
		case num, _ := <-fibchan:
			cnt, _ := <-countchan
			go printer(cnt, num, countchan, printchan)
			printcnt := <-printchan
			if cnt == amount && printcnt == amount {
				close(fibchan)
				close(countchan)
				close(printchan)
				stahp = true
			}
		default:
		}
	}
}
