package main

import (
    "fmt"
    "flag"
)

//    fmt.Printf("%3d: %.0f\n", 1, fibs[0])

func main() {
    var num int
    flag.IntVar(&num, "n", 10, "Set the number of fibonacci numbers to crunch [10]")
    flag.Parse()
    var f1, f2 float64 = 0, 1
    fmt.Printf("%3d: %.0f\n", 1, f1)
    for i:=1;i<num;i+=1 {
       f1, f2 = f2, f1+f2
       fmt.Printf("%3d: %.0f\n", i+1, f1)
    }
}
