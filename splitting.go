package main

import "fmt"

func main() {
	fmt.Print("Type shit: ")
	rtfm := make([]string, 99)
	for n := 0; n<20; n+=1 {
		fmt.Scan(&rtfm[n])
		fmt.Println("Current: ", rtfm[n])
		fmt.Println("Slices: ", rtfm[:n])
	}
}
