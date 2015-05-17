package main

import (
	"fmt"
)

/* test program to illustrate the go generate tool */

/* the binary generated here will be slower than go build, almost guaranteed. */
//go:generate go tool 6g -L -% -complete -race -o generator.6 generator.go
//go:generate go tool 6l -1 -8 -race -o generator generator.6
//go:generate rm generator.6
func main() {
	fmt.Print("Bark bark.\n")

	fmt.Print("Worden\n")

}
