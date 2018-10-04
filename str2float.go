package main

import (
	"fmt"
)

func main() {
	str := "-12.364"
	var n float64
	neg := false
	postdec := false
	// scalar for post-decimal
	decz := 0.10
	
	fmt.Println(str)
	
	if str[0] == '-' {
		neg = true
		str = str[1:]
	}

	for _, r := range str {
		if r == '.' {
			postdec = true
			continue
		}
		
		rN := float64(r - '0')
		fmt.Println(rN)
		
		if postdec {
			rN *= decz
			decz *= 0.1
			n += rN
		} else {
			n *= 10
			n += rN
		}
	}
	
	if neg {
		n *= -1
	}
	
	fmt.Println(n)
}

