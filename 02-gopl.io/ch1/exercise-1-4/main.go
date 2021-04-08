package main

import (
	"bufio"
	"fmt"
	"os"
)

type duplicate struct {
	count    int
	fileName string
}

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs
func main() {
	counts := make(map[string]duplicate)
	files := os.Args[1:]
	fmt.Println(files)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%s\t%d\t%s\n", n.fileName, n.count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]duplicate) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		dup := counts[input.Text()]
		dup.count++
		dup.fileName = f.Name()

		counts[input.Text()] = dup
	}
	// NOTE: ignoring potential errors from input.Err()
}
