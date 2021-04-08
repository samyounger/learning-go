package main

import "fmt"

func main() {
	// ==
	// <=
	// >=
	// !=
	// <
	// >
	x := 42 == 24
	y := 42 <= 53
	z := 42 >= 53
	a := 42 != 24
	b := 42 < 24
	c := 42 > 24

	fmt.Println(x, y, z, a, b, c)
}
