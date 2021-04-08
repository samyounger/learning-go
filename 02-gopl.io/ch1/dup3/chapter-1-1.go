package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Book optimised version
	fmt.Println(strings.Join(os.Args[1:], " "))

	// My version, including executing arg
	fmt.Println(strings.Join(os.Args[0:], " "))

	// My version, includes index of arg
	for i, a := range os.Args {
		fmt.Printf("Index: %v, Arg: %v\n", i, a)
	}
}
