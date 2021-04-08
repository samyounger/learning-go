package main

import "fmt"

type newInt int

var x newInt

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
