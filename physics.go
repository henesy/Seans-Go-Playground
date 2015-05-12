package main

import (
    "fmt"
    "math"
sc  "strconv"
)

type Direction int
const (
    UP Direction = iota
    DOWN
    LEFT
    RIGHT
)

type Nan float64
const (
    NAN Nan = iota
    UNKNOWN
)

type Physical struct {
    vi float64 /* initial vertical velocity (m/s)*/
    vf float64 /* final vertical velocity (m/s)*/
    a float64 /* vertical acceleration (m/s^2)*/
    d float64 /* vertical displacement (m)*/
    t float64 /* time (s)*/
    vhi float64 /* initial horizontal velocity (m/s)*/
    vhf float64 /* final horizontal velocity (m/s)*/
    ah float64 /* horizontal acceleration (m/s^2)*/
    fn float64 /* force normal (N)*/
    fg float64 /* force of gravity (N)*/
    fnet float64 /* net of forces (N)*/
    x float64 /* x force factor if angled, also Fm most likely (N)*/
    y float64 /* y force factor if angled, also Fg most likely (N)*/
    ang float64 /* angle of object (º)*/
    mass float64 /* mass of object (kg)*/
    pe float64 /* potential energy (Joules)*/
    ke float64 /* kinetic energy (Joules)*/
    dir Direction /* type Direction directional placeholder */
}

type Surface struct {
    mu float64 /* the µ coefficient of friction upon a surface */
    dis float64 /* distance or length of surface */
}

/* Objects is the interface for a Physical */
type Object interface {
    DeeEquals()
}

/* d=vi*t + (1/2)*a*t^2 */
func (p *Physical) DeeEquals() {
    p.d = p.vi*p.t + 0.5*p.a*(math.Pow(p.t, 2))
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

// (&thing[i].thing).method()

func main() {
    fmt.Print("\nPlease fill out a number for each value prompted for. If unknown, type \"?\".\n")
    fmt.Print("\nNumber of objects involved?: ")
    readin := ""
    fmt.Scanln(&readin)
    num, err := sc.Atoi(readin)
    check(err)
    fmt.Print(num, " objects.\n")
    objects := make([]Physical, num)

    for i:=0;i<num;i+=1 {
        words:=""
        var err error
        fmt.Print("Object ", i+1, ". \n\n")

        fmt.Print("Initial Vertical Velocity (meters/second): ")
        fmt.Scanln(&words)
        objects[i].vi, err = sc.ParseFloat(words, 64)
        fmt.Print(err, "\n")

        fmt.Print("Final Vertical Velocity (m/s): ")
        fmt.Scanln(&words)
        objects[i].vf, err = sc.ParseFloat(words, 64)
        fmt.Print(err, "\n")

        fmt.Print("Acceleration (meters/second^2): ")
        fmt.Scanln(&words)
        objects[i].a, err = sc.ParseFloat(words, 64)
        fmt.Print(err, "\n")

        fmt.Print("Displacement (meters): ")
        fmt.Scanln(&words)
        objects[i].d, err = sc.ParseFloat(words, 64)
        fmt.Print(err, "\n")

        fmt.Print("Time (seconds): ")
        fmt.Scanln(&words)
        objects[i].t, err = sc.ParseFloat(words, 64)
        fmt.Print(err, "\n")

        (&objects[i]).DeeEquals()
        fmt.Print("\n'd' equals: ", objects[i].d, "\n")
    }

}
