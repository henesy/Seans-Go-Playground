package main

import (
    "fmt"
   	"flag"
    "math"
    "math/big"
)

type op int
const (
    RUN op = iota
    DRUN
)
type action int
const (
    ADD action = iota
    SUB
)

/*
solves for a flag-able number of digits of pi using the Nilakantha series:
π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12)...
*/

/* calculates for pi */
func calcPi(pi, a, b, c *big.Float, todo *op, act *action) {
	f := big.NewFloat(4)
    //tmp := big.NewFloat(0)
	tot1 := big.NewFloat(0)
    tot2 := big.NewFloat(0)
    frac := big.NewFloat(0)
    oldpi := big.NewFloat(0)
    oldpi.Set(pi)
    //n1 := big.NewFloat(-1)

    tot1.Add(a, b)
    tot2.Add(tot1, c)
    frac.Quo(f, tot2)

    if(*act == ADD) {
        pi.Add(oldpi, frac)
    } else {
        pi.Sub(oldpi, frac)
    }
    
    *act = action(int(math.Abs(float64(int(*act) - 1))))
    *todo = DRUN
}


/* concurrent pi calculator */
func main() {
    var iterations uint64
    flag.Uint64Var(&iterations, "n", 100000, "Set the number of iterations to perform [100,000]")
	flag.Parse()
	var todo op
    act := ADD
    pi := big.NewFloat(3)
	a := big.NewFloat(2)
	b := big.NewFloat(3)
	c := big.NewFloat(4)
	//f := big.NewFloat(4)
    //var actnext act = ADD
    tmp := big.NewFloat(0)
    two := big.NewFloat(2)
    //n1 := big.NewFloat(-1)
    //f1 := big.NewFloat(4)
    //f.Mul(f1, n1)
	for i:=0;uint64(i) < iterations;i+=1 {
        todo = RUN
        //go calcPi(pi, a, b, c, &todo, &act)
        olda := big.NewFloat(0)
        oldb := big.NewFloat(0)
        oldc := big.NewFloat(0)
        olda.Set(a)
        oldb.Set(b)
        oldc.Set(c)
        go calcPi(pi, a, b, c, &todo, &act)
        for(todo == RUN) {}
        a = tmp.Add(olda, two)
        b = tmp.Add(oldb, two)
        c = tmp.Add(oldc, two)
    }
    v, _ := pi.Float64()
	fmt.Print(v, "\n")
    
}





