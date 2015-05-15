package main

/*
#include <stdio.h>
*/
import "C"
import (
    "fmt"
sc  "strconv"
    "unicode/utf8"
)

func main() {
    e := C.getchar()
    f := int(C.getchar())
    fmt.Print(f)
    g := sc.Itoa(f)
    h, _ := utf8.DecodeRuneInString(g)
    fmt.Printf(" or %c or %s\n", h, g)
    
    C.printf("%c\n", C.int(e))
}

