package main

import (
    "fmt"
)

type direction int
const (
    UP direction = iota
    DOWN 
    LEFT 
    RIGHT 
    NONE 
)

func printArr(arr [][]uint64) {
    
    for i:=0;i<len(arr);i++ {
        fmt.Println(arr[i])
    }
    fmt.Print(": ")
}


func main() {
    arr := make([][]uint64, 4)
    for i := range arr {
        arr[i] = make([]uint64, 4)
    }
    arr[0][2], arr[3][1] = 2, 2
    printArr(arr)

    game := true
    for game == true {
        dir := NONE
        str := ""
        fmt.Scanln(&str)
        switch str {
            default: dir = NONE
            case "w": dir = UP
            case "a": dir = LEFT
            case "s": dir = DOWN
            case "d": dir = RIGHT
        }
        



        printArr(arr)
    }

}

