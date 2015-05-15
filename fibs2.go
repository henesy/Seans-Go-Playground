package main

import "fmt"

func fib(num int, fibchan chan int) {

}

func main() {
    fibchan := make(chan float32, 100)
    go fib()
}
