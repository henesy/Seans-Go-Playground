package main

import (
    "fmt"
    "flag"
    "math/big"
)

/* basic single-threaded uncapped fibonacci. Only prints once. */

func main() {
    var num int
    flag.IntVar(&num, "n", 10, "Set the number of fibonacci numbers to crunch [10]")
    flag.Parse()
    f1, f2 := big.NewInt(1), big.NewInt(1)
    for i:=1;i<num+1;i+=1 {
       tmp := new(big.Int)
       tmp.Add(f1, f2)
       f1 , f2 = f2, tmp
    }
    fmt.Printf("%s\n", f1.String())
}
