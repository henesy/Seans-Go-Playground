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

type Num float64
const (
    NAN Num = iota
    UNKNOWN
    KNOWN
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

/* counterpart to Physical that shows if we know a value or not (have solved for, etc.) */
type Knowledge struct {
    vi Num /* initial vertical velocity (m/s)*/
    vf Num /* final vertical velocity (m/s)*/
    a Num /* vertical acceleration (m/s^2)*/
    d Num /* vertical displacement (m)*/
    t Num /* time (s)*/
    vhi Num /* initial horizontal velocity (m/s)*/
    vhf Num /* final horizontal velocity (m/s)*/
    ah Num /* horizontal acceleration (m/s^2)*/
    fn Num /* force normal (N)*/
    fg Num /* force of gravity (N)*/
    fnet Num /* net of forces (N)*/
    x Num /* x force factor if angled, also Fm most likely (N)*/
    y Num /* y force factor if angled, also Fg most likely (N)*/
    ang Num /* angle of object (º)*/
    mass Num /* mass of object (kg)*/
    pe Num /* potential energy (Joules)*/
    ke Num /* kinetic energy (Joules)*/
    dir Num /* type Direction directional placeholder */
}

/* surfaces to be calculated against */
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

/* check for error, panic if found */
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
    knowns := make([]Knowledge, num)

    for i:=0;i<num;i+=1 {
        words:=""
        var err error
        fmt.Print("Object ", i+1, ". \n\n")

        fmt.Print("Initial Vertical Velocity (meters/second): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].vi, err = sc.ParseFloat(words, 64)
            knowns[i].vf=KNOWN
            fmt.Print(err, "\n")
            check(err)
            knowns[i].vf=KNOWN
        } else if words == "?" {
            knowns[i].vf=UNKNOWN
        }

        fmt.Print("Final Vertical Velocity (m/s): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].vf, err = sc.ParseFloat(words, 64)
            fmt.Print(err, "\n")
            check(err)
            knowns[i].vf=KNOWN
        } else if words == "?" {
            knowns[i].vf=UNKNOWN
        }

        fmt.Print("Acceleration (meters/second^2): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].a, err = sc.ParseFloat(words, 64)
            fmt.Print(err, "\n")
            check(err)
            knowns[i].a=KNOWN
        } else if words == "?" {
            knowns[i].a=UNKNOWN
        }

        fmt.Print("Displacement (meters): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].d, err = sc.ParseFloat(words, 64)
            fmt.Print(err, "\n")
            check(err)
            knowns[i].d=KNOWN
        } else if words == "?" {
            knowns[i].d=UNKNOWN
        }

        fmt.Print("Time (seconds): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].t, err = sc.ParseFloat(words, 64)
            fmt.Print(err, "\n")
            check(err)
            knowns[i].t=KNOWN
        } else if words == "?" {
            knowns[i].t=UNKNOWN
        }

        (&objects[i]).DeeEquals()
        fmt.Print("\n'd' equals: ", objects[i].d, "\n")
    }

}
