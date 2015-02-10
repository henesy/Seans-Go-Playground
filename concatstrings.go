package main

import "fmt"

func main() {
	word1 := "/path/to"
	word2 := "/acmd"
	word3 := word1 + word2
	fmt.Print(word3, "\n", word2 + word1 + "/anewcmd\n")
}
