package main

import "fmt"

func main() {
	f := foo(1, 1)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo(a int, b int) func() string {
	sum := a + b
	return func() string {
		sum++
		return fmt.Sprintf("The sum is %v", sum)
	}
}
