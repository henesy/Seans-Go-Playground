package main

// #include <stdio.h>
//// #include <stdlib.h>
import "C"
import (
    "fmt"
//    "unsafe"
)

func main() {
    //var e *[20]C.char
    //C.scanf("%s", unsafe.Pointer(&e[0]) )
    f := C.getchar()
    //g := C.GoString(e)
    fmt.Println(f)
    //C.free(unsafe.Pointer(e))
}

