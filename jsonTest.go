package main

import (
	"fmt"
	"encoding/json"
	"os"
	"flag"
)

type Sprite struct {
	X		int
	Y		int
	Name	string
}

/* checks errors */
func check(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}


/* encodes and/or reads a struct-arr.Y from a json-encoded file */
func main() {
	var encode bool
	flag.BoolVar(&encode, "encode", true, "Will encode if true, decode if false.")
	flag.Parse()

	//populate the arr.Y
	arr := make([]Sprite, 5)
	i := 0
	for p, _ := range arr {
		arr[p].X = i
		arr[p].Y = i+1
		i++
		fmt.Printf("Name for sprite %d: ", p)
		rsps := ""
		fmt.Scanln(&rsps)
		arr[p].Name = rsps
	}

	//encode or decode via flag decision
	if encode == true {
		file, err := os.OpenFile("jsontest_writeout.txt", os.O_WRONLY, 0644)
		check(err)
		enc := json.NewEncoder(file)
		check(err)
		err = enc.Encode(arr)
		fmt.Printf("%v\n", arr)
		check(err)
	} else {
		file, err := os.OpenFile("jsontest_writeout.txt", os.O_RDONLY, 0644)
		check(err)
		dec := json.NewDecoder(file)
		newArr := make([]Sprite, 5)
		for i := 0;dec.More() && i < len(newArr);i++ {
			dec.Decode(&newArr[i])
			check(err)
		}
		fmt.Printf("%v\n", newArr)
	}
}

