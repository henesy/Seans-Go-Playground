package main

import (
    "fmt"
sc  "strconv"
)

type maf func(...int)(int)

func decide(nums ...int) maf {
    tot := 0
    for _, i := range nums {
        tot += i
    }

    if tot < 100 {
        return func(numz ...int)(int) {
            sum := 0
            for _, n := range numz {
                sum += n
            }
            return sum
        }
    }

    return func(numz ...int)(int) {
        dif := 0
        for _, n := range numz {
            dif -= n
        }
        return dif
    }
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

/* adds all integers < 100; subs >= 100 */

func main() {
    words := ""
    fmt.Print("How many integers?: ")
    fmt.Scanln(&words)
    nw, err := sc.Atoi(words)
    check(err)
    arr := make([]int, nw)

    /* the index value is the default value -- the first */
    for pos := range arr {
        wrd := ""
        fmt.Print("integer: ")
        fmt.Scanln(&wrd)
        n, err := sc.Atoi(wrd)
        check(err)
        arr[pos] = n
    }

    mathing := decide(arr...)
    result := mathing(arr...)
    fmt.Print(result, "\n")
}
