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
    DEquals()
}

/* */

/* d=vi*t + (1/2)*a*t^2 */
func (p *Physical) DEquals() {
    p.d = p.vi*p.t + 0.5*p.a*(math.Pow(p.t, 2))
}

/* vf^2 = vi^2 + 2*a*d */
func (p *Physical) VFSquaredEquals() {
    p.vf = math.Sqrt(math.Pow(p.vi, 2)+2*p.a*p.d)
}

/* t = (vf - vi)/a */
func (p *Physical) TEquals() {
    p.t = (p.vf - p.vi)/p.a
}

/* check for error, panic if found */
func check(err error) {
    if err != nil {
        panic(err)
    }
}

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
        fmt.Print("\nObject ", i+1, ". \n")

        fmt.Print("Initial Vertical Velocity (meters/second): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].vi, err = sc.ParseFloat(words, 64)
            knowns[i].vf=KNOWN
            check(err)
            knowns[i].vi=KNOWN
        } else if words == "?" || words == "" || words == " " {
            knowns[i].vi=UNKNOWN
        }

        fmt.Print("Final Vertical Velocity (m/s): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].vf, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].vf=KNOWN
        } else if words == "?" || words == "" || words == " " {
            knowns[i].vf=UNKNOWN
        }

        fmt.Print("Acceleration (meters/second^2): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].a, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].a=KNOWN
        } else if words == "?" || words == "" || words == " " {
            knowns[i].a=UNKNOWN
        }

        fmt.Print("Displacement (meters): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].d, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].d=KNOWN
        } else if words == "?" || words == "" || words == " " {
            knowns[i].d=UNKNOWN
        }

        fmt.Print("Time (seconds): ")
        fmt.Scanln(&words)
        if words != "?" {
            objects[i].t, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].t=KNOWN
        } else if words == "?" || words == "" || words == " " {
            knowns[i].t=UNKNOWN
        }

        /* perform calculations to solve for a missing variable; must know at least 3 things */
        fmt.Print("\nResults: \n")
        if knowns[i].t == KNOWN && knowns[i].a == KNOWN && knowns[i].vi == KNOWN {
            (&objects[i]).DEquals()
            fmt.Print("'d' = ", objects[i].d, "\n")
        }
        if knowns[i].vi == KNOWN && knowns[i].a == KNOWN && knowns[i].d == KNOWN {
            (&objects[i]).VFSquaredEquals()
            fmt.Print("'vf' = ", objects[i].vf, "\n")
        }
        if knowns[i].vi == KNOWN && knowns[i].vf == KNOWN && knowns[i].a == KNOWN {
            (&objects[i]).TEquals()
            fmt.Print("'t' = ", objects[i].t, "\n")
        }

    }

}
