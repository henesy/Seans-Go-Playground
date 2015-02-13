package main
import (
"fmt"
)
func main() {
    in := 4 //factorial number
    nin := in
    for in > 1 {
        in = (in-1)
        nin = (nin * in)
    }
    fmt.Print("Factorial is: ", nin, "\n")
}
