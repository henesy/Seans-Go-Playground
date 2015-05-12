package main

import (
    "fmt"
)

/*
<
•skelterjohn>
lungaro: var blah *T; blah = new(T) ... *blah = &T{}

func (p *Physical) DEquals(knowns *[]Knowledge) {
    p.d = p.vi*p.t + 0.5*p.a*(math.Pow(p.t, 2))
    (*knowns)[p.i].d = KNOWN
}
*/

func main() {
    arr := make([]int, 10)
    var ß *int
    //var µ **int
    nums := &arr
    a := 1
    b := 2
    c := 3
    x := &a
    y := &b
    z := &c
    ß = &b
    //**µ = *ß + a


    (*nums)[0] = a
    fmt.Print(nums, *nums, &nums, "\n")
    fmt.Print(arr[0], "\n")
    fmt.Print(a, b, c, *x, y, *z, ß, "\n")
}
