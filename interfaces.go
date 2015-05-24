package main

import (
    "fmt"
    "math"
)

/* takes structs in to calculate certain aspects; can be 1d-3d */
type Calculator interface {
    Area() float64
    Circumference() float64
    Volume() float64
    Density() float64
}

/* deals with slices of ints and shifts their contents */
type Adjuster interface {
    ShiftLeft() []int
    ShiftRight() []int
    Flip() []int
    Remix() []int
}

/* types must be named; []int and []float64 are unnamed types, thus we name them */
type Iarr []int
type Farr []float64

/* exact defintions of circle may be 1d-3d; sphere or circle */
type Circle struct {
    radius float64
    mass float64
}

type Square struct {
    length float64
    mass float64
}

/* slices to play with */
var arr1 = []int{5, 15, 7, 2, 8, 9}
var arr2 = []int{2, 6}
var arr3 = []int{19, 65, 2, 4, 6, 8, 22}
var arr4 = []float64{76.25, 15235.15, 8436.53, 9240.124, 241.24}

/* methods to calculate aspects of shapes */
/* square methods */
func (s Square) Area() float64 {
    return math.Pow(s.length, 2)
}

func (s Square) Circumference() float64 {
    return s.length * 4
}

func (s Square) Volume() float64 {
    return math.Pow(s.length, 3)
}

func (s Square) Density() float64 {
    ca := Calculator(s)
    return s.mass / ca.Volume()
}

/* circle methods */
func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Circumference() float64 {
    return math.Pi * 2 * c.radius
}

func (c Circle) Volume() float64 {
    return (4.0/3.0) * math.Pi * math.Pow(c.radius, 3)
}

func (c Circle) Density() float64 {
    ca := Calculator(c)
    return c.mass / ca.Volume()
}

/* methods to adjust slices or arrays */
func (a Iarr) ShiftRight() Iarr {
    newArr := make(Iarr, len(a))
    for i := 0; i < len(a)-1; i++ {
        newArr[i+1] = a[i]
    }
    return newArr
}

func (a Iarr) ShiftLeft() Iarr {
    newArr := make(Iarr, len(a))
    for i := len(a)-1; i > 0; i-- {
        newArr[i-1] = a[i]
    }
    return newArr
}

func (a Iarr) Flip() Iarr {
    newArr := make(Iarr, len(a))
    h := len(a)-1
    for i := 0; i < len(a); i++ {
            newArr[i] = a[h]
            fmt.Println(h)
            h--
    }
    return newArr
}

func (a Iarr) Remix() Iarr {
    // bleh, not implemented, yet
    return a
}

func (a Farr) ShiftLeft() Farr {
    newArr := make(Farr, len(a))
    for i := len(a)-1; i > 0; i-- {
        newArr[i-1] = a[i]
    }
    return newArr
}

func (a Farr) ShiftRight() Farr {
    newArr := make(Farr, len(a))
    for i := 0; i < len(a)-1; i++ {
        newArr[i+1] = a[i]
    }
    return newArr
}

func (a Farr) Flip() Farr {
    newArr := make(Farr, len(a))
    h := len(a)-1
    for i := 0; i < len(a); i++ {
            newArr[i] = a[h]
            h--

    return newArr
}

func (a Farr) Remix() Farr {
    //not implemented either
    return a
}


/* quick demo program to test and toy with interfaces */

func main() {
    fmt.Print(arr1, arr2, arr3, arr4, "\n")
    arr1 = []int(Iarr(arr1).ShiftLeft())
    arr2 = []int(Iarr(arr2).ShiftRight())
    arr3 = []int(Iarr(arr3).Flip())
    arr4 = []float64(Farr(arr4).ShiftLeft())
    arr4 = []float64(Farr(arr4).Flip())
    fmt.Print(arr1, arr2, arr3, arr4, "\n")
    c := Circle{4, 6}
    s := Square{8, 3}
    fmt.Print(c, s, "\n")
    ac := c.Area()
    as := s.Area()
    cc := c.Circumference()
    sd := s.Density()
    fmt.Print(ac, as, cc, sd, "\n")

}
