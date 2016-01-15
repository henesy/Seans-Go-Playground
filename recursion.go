package main

import (
	"fmt"
	"os"
sc	"strconv"
)

/* recursive function that checks size */
func listHandler(lst []int, top int) {
	if len(lst) < 1{
		fmt.Printf("The largest number is: %d\n", top)
	} else if lst[0] > top {
		listHandler(lst[1:], lst[0])
	} else {
		listHandler(lst[1:], top)
	}
}


/* Recursive largest num in list calculator */
func main() {
	lst := make([]int, 0, len(os.Args))
	for _, n := range os.Args[1:] {
		i, _ := sc.Atoi(n)
		lst = append(lst, i)
	}

	listHandler(lst, lst[0])
}

