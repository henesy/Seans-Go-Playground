package main

import (
    "fmt"
    "flag"
    "math/big"
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

    pi := big.NewRat(3, 1)

    var a, b, c int64 = 2, 3, 4
    var actnext act = ADD
    temp0 := new(big.Rat)
    temp0.Set(pi)
    for i:=0;int64(i) < num;i+=1 {
        temp1, temp2, temp3, temp4 := big.NewRat(4, 1), big.NewRat(1, (a*b*c)), new(big.Rat), new(big.Rat)
        /*
        temp1 is the 4 of the 4/x sequence
        temp2 is the x of the " to be
        */
        temp3.Mul(temp1, temp2)
        if actnext == ADD {
            temp4.Add(temp0, temp3)
            actnext = SUB
        } else if actnext == SUB {
            temp4.Sub(temp0, temp3)
            actnext = ADD
        }
        temp0.Set(temp4)

        a+=2
        b+=2
        c+=2
    }
    pi.Set(temp0)
    z, _ := pi.Float64()
    fmt.Print(z, "\n")
}
