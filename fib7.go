package main

import (
    "fmt"
    "flag"
)

/* basic single-threaded capped fibonacci using unsigned integers */

func main() {
    var num uint64
    flag.Uint64Var(&num, "n", 10, "Set the number of fibonacci numbers to crunch [10]")
    flag.Parse()
    var f1, f2 uint64 = 0, 1
    fmt.Print(1, ": ", f1, "\n")
    var i uint64
    for i=1;i<num;i+=1 {
       f1, f2 = f2, f1+f2
       fmt.Print(i+1, ": ", f1, "\n")
    }
}
