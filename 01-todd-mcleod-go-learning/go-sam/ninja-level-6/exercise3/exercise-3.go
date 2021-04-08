package main

import "fmt"

func main() {
	defer foo()

	fmt.Println("This Prints First")
}

func foo() int {
	fmt.Println("This Prints Second")
	return 5
}
