package main

import (
	"fmt"
	"github.com/derekparker/trie"
)

type File struct {
	qid			uint
	dir			bool
	contents	[]byte
}

/* Tests the usage of a trie in an fs-like manner */
func main() {
	fmt.Println("Initializing trie…")
	t := trie.New()

	fmt.Println("Adding root node")
	var root File
	root.qid = 0
	root.dir = true
	t.Add("/", root)

	fmt.Println("Reading root node")
	node, ok := t.Find("/")
	if ok {
		// Read contents of file
		file := node.Meta().(File)
		fmt.Println(file)
	}

	fmt.Println("Adding some subdirs…")
	t.Add("/sys", 1)
	usr := t.Add("/usr", 2)
	t.Add("/tmp", 3)

	// Showing all nodes with a "/" prefix
	fmt.Println(t.PrefixSearch("/"))

	fmt.Println("Getting children of /")
	fmt.Println(node.Children())

	fmt.Println("Getting children of /usr")
	fmt.Println(usr.Children())
}

