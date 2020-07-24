package main

import (
	"flag"
	"fmt"
	"crypto/sha256"
	"path/filepath"
	"os"
	"io"
	"io/ioutil"
	"strings"
)

// String pair
type Pair struct {
	dupe	string	// Duplicate file path
	of		string	// What dupe is a duplicate of
}

var (
	chatty		bool				// Verbose output?
	noRecurse	bool				// Do we recurse child directories?
	noFilter	bool				// Skip names like .git?
	root		string				// Root directory to begin with
	maxFiles	uint64				// Maximum files which can be deleted

	files		map[string]string	// File name → hash mapping
	// Path elements which should constitute filtering out a path
	undesirables = []string{
					".git/",
					".hg/",
					".svn/",
				}
)


// Determine whether to skip/filter out a path - this is bad and slow
func bad(path string) bool {
	if noFilter {
		return false
	}

	for _, bad := range undesirables {
		if strings.Contains(path, bad) {
			return true
		}
	}

	return false
}

// Generate hash string from file path
func hashify(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return string(hash.Sum(nil)), nil
}

// Walk a file
func walker(path string, info os.FileInfo, err error) error {
	// Skip parsing directories
	if info.IsDir() || bad(path) {
		if chatty {
			fmt.Fprintln(os.Stderr, "skipping: ", path)
		}

		return nil
	}

	if chatty {
		fmt.Fprintln(os.Stderr, "ingesting: ", path)
	}

	hash, err := hashify(path)
	if err != nil {
		return err
	}

	files[path] = hash

	return nil
}


// De-duplicate files in a directory based on sha256 hash
func main() {
	flag.BoolVar(&chatty, "D", false, "Verbose debug output")
	flag.BoolVar(&noRecurse, "R", false, "Do not recurse through child directories")
	flag.BoolVar(&noFilter, "F", false, "Recurse on problematic file paths like .git")
	flag.StringVar(&root, "d", "./", "Directory to de-duplicate inside")
	flag.Uint64Var(&maxFiles, "m", 4096, "Maximum number of files which can be deleted")
	flag.Parse()

	files = make(map[string]string)

	if !noRecurse {
		// Recurse
		err := filepath.Walk(root, walker)
		if err != nil {
			fmt.Fprintln(os.Stderr, "err: cannot walk -", err)
			os.Exit(1)
		}
	} else {
		// No recursion
		contents, err := ioutil.ReadDir(root)
		if err != nil {
			fmt.Fprintln(os.Stderr, "err: cannot read dir -", err)
			os.Exit(2)
		}

		for _, info := range contents {
			// Skip directories
			if info.IsDir() {
				continue
			}

			path := info.Name()
			hash, err := hashify(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "err: could not hash -", err)
				os.Exit(3)
			}

			files[path] = hash
		}
	}

	// Set up for reverse lookup of hash→path
	reversed := make(map[string]string)
	dupes := make([]Pair, 0, maxFiles)

	for path, hash := range files {
		if of, ok := reversed[hash]; ok {
			// This is a duplicate
			dupes = append(dupes, Pair{path, of})

			continue
		}

		reversed[hash] = path
	}

	// Inform the user
	fmt.Println("KEEPING:")
	for _, path := range reversed {
		fmt.Println("	", path)
	}

	fmt.Println("DELETING:")
	for _, pair := range dupes {
		fmt.Printf("\t%v\t→\t%v\n", pair.dupe, pair.of)
	}

	// Confirm with the user
	var in string
	fmt.Print("Delete duplicates? [y/n]: ")
	fmt.Scanln(&in)

	if in != "y" {
		fmt.Println("Not deleting, exiting cleanly.")
		return
	}

	fmt.Printf("Ok, deleting %v files.\n", len(dupes))

	// Delete
	for _, pair := range dupes {
		path := pair.dupe
		err := os.Remove(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, `err: could not delete "%v" - %v\n`, path, err)
			fmt.Fprintln(os.Stderr, "Stopping.")
			os.Exit(4)
		}
	}

	fmt.Println("Done.")
}

