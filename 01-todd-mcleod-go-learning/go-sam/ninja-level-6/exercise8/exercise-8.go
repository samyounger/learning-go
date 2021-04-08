package main

import "fmt"

func main() {
	f := foo()
	fmt.Println("Before calling func second time")
	fmt.Println(f())
}

func foo() func() string {
	return func() string {
		return "Inside Func"
	}
}
