package main

import "fmt"

//fmt.Printf("%3d: %.0f\n", i+1, fibs[i])
var count int = 1

func printer(num1 int, num2 float64, countchan chan int, printchan chan int) {
       fmt.Printf("%3.0d: %.0f\n", num1, num2)
       printchan<-num1
}

func fib(fibchan chan float64, countchan chan int, fibsize int) {
    fibs := make([]float64, fibsize)
    fibs[0],fibs[1] = 0,1
    fibchan <- fibs[0]
    countchan <- count
    fibchan <- fibs[1]
    count+=1
    countchan <- count

    for i:=2;i<len(fibs);i+=1 {
        count=i+1
        nums := fibs[i-2:i]
        fibs[i] = nums[0] + nums[1]
        countchan <- count
        fibchan <- fibs[i]
    }
}

func main() {
    amount:=0
    fmt.Print("Crunch how many fibonacci numbers?: ")
    fmt.Scanln(&amount)
    fibchan := make(chan float64, amount)
    countchan := make(chan int, 3)
    printchan := make(chan int, amount)
    go fib(fibchan, countchan, amount)
    var cnt int
    for stahp:=false;(!stahp); {
        select {
        case num, _ := <-fibchan:
                cnt, _ = <-countchan
                go printer(cnt, num, countchan, printchan)
                printcnt:=<-printchan
                if cnt == amount && printcnt == amount {
                    close(fibchan)
                    close(countchan)
                    close(printchan)
                    stahp=true
                }
        default:
        }
    }
}
