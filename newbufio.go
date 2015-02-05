package main

import (
"fmt"
"bufio"
"os"
"strings"
"os/exec"
"syscall"
)

func registrar() {
	var stringz string
	reader := os.Stdin
	env := os.Environ()

	sc := bufio.NewScanner(reader)

	scanerr := sc.Scan() // get input
	if scanerr == false {
		fmt.Println("Whoops! Scanner broke!")
	}
	stringz = sc.Text() // store input

	newstringz := make([]string, 99)

	/* split the user's input into a command and arguments */
	newstringz = strings.SplitN(stringz, " ", 2)

	app := newstringz[0]

	if newstringz[1] == "" {
		newstringz[1] = " "
	} else if len(newstringz) == 1 {
		newstringz[1] = " "
	} 

	/* split arguments into a slice (required by syscall.Exec) and give it a nice name*/
	args := strings.Split(newstringz[1], " ")
	fmt.Println(args)

	binary, err := exec.LookPath(app) // get path for program (`/bin/ls` for example)
 	   if err != nil {
 	       panic(err)
 	   }

	fmt.Println("DEBUG: ", app, args)

	execErr := syscall.Exec(binary, args[0:], env)
	    if execErr != nil {
	        panic(execErr)
 	   }
}

func main() {

	for 1==1 {
	fmt.Print("SVI% ")
	registrar()

	}
} // main end
