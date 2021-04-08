package main

import "fmt"

func main() {
	x := func() string {
		return "Here in first func"
	}
	f := foo(x)
	fmt.Println(f)
}

func foo(f func() string) string {
	fmt.Println("Here in second func")
	return f()
}
