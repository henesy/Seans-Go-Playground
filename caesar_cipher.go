package main

import (
    "fmt"
    "flag"
)



func main() {
	var cipher, dictName string
	var timeToRun uint64
	flag.StringVar(&cipher, "c", "", "Cipher to bruteforce [Prompt]")
	flag.StringVar(&dictName, "d", "", "Dictionary to use [None]")
	flag.Uint64Var(&timeToRun, "t", 0, "Time limit to run to [0]")
	flag.Parse()
	fmt.Println("Warning: A caesar cipher has a potential 26^26 permutations, this operation takes awhile, let alone any dictionary tests that one might wish to run. Expect long run times. No complaining allowed.")
	
	
	



}

