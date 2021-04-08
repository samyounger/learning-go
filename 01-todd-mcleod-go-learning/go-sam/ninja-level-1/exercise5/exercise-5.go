package main

import "fmt"

type foo int

var x foo = 42
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 24
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
