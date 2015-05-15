package main

import (
    "fmt"
    "math/rand"
)

/* adds all integers < 100; subs >= 100 */

func main() {
    rng := rand.New(rand.NewSource(12412509358))
    rnd := rng.Int()

    // value of return
    message := func() string {
        txt := ""
        if rnd > 100 && rnd < 1000 {
            txt = "More than 100 and less than 1000!"
        } else if rnd < 100 && rnd > 10 {
            txt = "Less than 100 and greater than 10!"
        } else if rnd < 10 {
            txt = "Less than 10!"
        } else {
            txt = "IT'S OVER 1000!!¡¡!!¡¡!!"
        }
        return txt
    }() /* this is why the value is returned, the () operates var as a func() */

    // function
    printgen := func(words string) func() {
        return func() {
            fmt.Println(words)
        }
    } /* lack of the operative () after closure, thus it inherits the func()
        definition */

    /* printgen returns a function that has the message already embedded */
    printer := printgen(message)
    printer()
}
