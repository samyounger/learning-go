package main

import "fmt"

func main() {
	f := foo()
	bi, bs := bar()

	fmt.Println(f)
	fmt.Println(bi, bs)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 6, "Here in bar"
}
