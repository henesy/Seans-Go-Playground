package main

import (
    "fmt"
    "math"
sc  "strconv"
    "flag"
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
    i int /* placeholder for array */
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
    w float64 /* Work (Joules)*/
    pow float64 /* Power (Watts)*/
    p float64 /* Momentum (Newtons?)*/
    dir Direction /* type Direction directional placeholder */
}

/* counterpart to Physical that shows if we know a value or not (have solved for, etc.) */
type Knowledge struct {
    i int /* placeholder for array */
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
    w Num
    pow Num
    p Num
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
    VFSquaredEquals()
    TEquals()
}

/* Functions to perform mathematical formulas */

/* d=vi*t + (1/2)*a*t^2 */
func (p *Physical) DEquals(knowns *[]Knowledge) {
    p.d = p.vi*p.t + 0.5*p.a*(math.Pow(p.t, 2))
    /* the []index needs to read from the knowns pointer, thus reify the pointer
    then write to index */
    (*knowns)[p.i].d = KNOWN
}

/* vf^2 = vi^2 + 2*a*d */
func (p *Physical) VFSquaredEquals(knowns *[]Knowledge) {
    p.vf = math.Sqrt(math.Pow(p.vi, 2)+2*p.a*p.d)
    (*knowns)[p.i].vf = KNOWN
}

/* t = (vf - vi)/a */
func (p *Physical) TEquals(knowns *[]Knowledge) {
    p.t = (p.vf - p.vi)/p.a
    (*knowns)[p.i].t = KNOWN
}

/* F = ma */
func (p *Physical) Weight(knowns *[]Knowledge) {
    p.fg = p.m * (-9.8)
    (*knowns)[p.i].fg = KNOWN
}

/* F = ma */
func (p *Physical) Fn(knowns *[]Knowledge) {
    p.fn = math.Abs(p.fg)
    (*knowns)[p.i].fn = KNOWN
}

/* a = F/m */


/* m = F/a */

/* W = F*d */

/* d = W/F */

/* P = m * (vf-vi) */

/* check for error, panic if found */
func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    var num int
    var err error
    var specify bool
    flag.IntVar(&num, "n", 1, "Number of objects to crunch for.")
    flag.BoolVar(&specify, "s", false, "Whether or not to choose the number of objects.")
    flag.Parse()
    fmt.Print("\nIf a value is unknown, type \"?\".\n")
    if readin:=""; specify == true {
        fmt.Print("\nNumber of objects involved?: ")
        fmt.Scanln(&readin)
        num, err = sc.Atoi(readin)
        check(err)
    }
    fmt.Print(num, " objects.\n")
    objects := make([]Physical, num)
    knowns := make([]Knowledge, num)

    for i:=0;i<num;i+=1 {
        words:=""
        fmt.Print("\nObject ", i+1, ". \n")
        objects[i].i, knowns[i].i = i, i

        fmt.Print("Initial Vertical Velocity (meters/second): ")
        fmt.Scanln(&words)
        if words == "?" || words == "" || words == " " {
            knowns[i].vi=UNKNOWN
        } else if words != "?" {
            objects[i].vi, err = sc.ParseFloat(words, 64)
            knowns[i].vf=KNOWN
            check(err)
            knowns[i].vi=KNOWN
        }


        fmt.Print("Final Vertical Velocity (m/s): ")
        fmt.Scanln(&words)
        if words == "?" || words == "" || words == " " {
            knowns[i].vf=UNKNOWN
        } else if words != "?" {
            objects[i].vf, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].vf=KNOWN
        }

        fmt.Print("Acceleration (meters/second^2): ")
        fmt.Scanln(&words)
        if words == "?" || words == "" || words == " " {
            knowns[i].a=UNKNOWN
        } else if words != "?" {
            objects[i].a, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].a=KNOWN
        }

        fmt.Print("Displacement (meters): ")
        fmt.Scanln(&words)
        if words == "?" || words == "" || words == " " {
            knowns[i].d=UNKNOWN
        } else if words != "?" {
            objects[i].d, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].d=KNOWN
        }

        fmt.Print("Time (seconds): ")
        fmt.Scanln(&words)
        if words == "?" || words == "" || words == " " {
            knowns[i].t=UNKNOWN
        } else if words != "?" {
            objects[i].t, err = sc.ParseFloat(words, 64)
            check(err)
            knowns[i].t=KNOWN
        }

        /* perform calculations to solve for a missing variable; must know at least 3 things */
        fmt.Print("\nResults: \n")
        if knowns[i].t == KNOWN && knowns[i].a == KNOWN && knowns[i].vi == KNOWN {
            (&objects[i]).DEquals(&knowns)
            fmt.Print("'d' = ", objects[i].d, "\n")
        }
        if knowns[i].vi == KNOWN && knowns[i].a == KNOWN && knowns[i].d == KNOWN {
            (&objects[i]).VFSquaredEquals(&knowns)
            fmt.Print("'vf' = ", objects[i].vf, "\n")
        }
        if knowns[i].vi == KNOWN && knowns[i].vf == KNOWN && knowns[i].a == KNOWN {
            (&objects[i]).TEquals(&knowns)
            fmt.Print("'t' = ", objects[i].t, "\n")
        }

    }

}
