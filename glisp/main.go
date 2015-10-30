package main

import (
	"fmt"
)



/* an attempt at a form of lisp REPL */
func main() {
	//var exec func() interface{}
	var usrin string
	prompt := "[%d]> "
	errs := 0


	for(usrin != "exit") {
		fmt.Printf(prompt, errs)	
		fmt.Scanln(&usrin)
			
	
	
	
	}
	
}


