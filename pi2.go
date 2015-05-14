package main

import (
    "fmt"
    "flag"
)

/*
solves for a flag-able number of digits of pi using the Nilakantha series:
π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12)...
*/
type act int
const (
    ADD act = iota
    SUB
)

func main() {
    var num int64
    flag.Int64Var(&num, "n", 100000, "Set the number of iterations to perform. [100,000]")
    flag.Parse()

    var pi float64 = 3.0

    var a, b, c float64 = 2.0, 3.0, 4.0
    var actnext act = ADD
    for i:=0;int64(i) < num;i+=1 {
        if actnext == ADD {
            pi += 4.0/(a*b*c)
            a+=2
            b+=2
            c+=2
            actnext = SUB
        } else if actnext == SUB {
            pi -= 4.0/(a*b*c)
            a+=2
            b+=2
            c+=2
            actnext = ADD
        }
    }
    fmt.Print(pi, "\n")
}
