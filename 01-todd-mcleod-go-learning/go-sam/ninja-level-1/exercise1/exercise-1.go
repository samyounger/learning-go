package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	// Format 1
	fmt.Println(x, y, z)

	// Format 2
	fmt.Printf("%v\n%v\n%v", x, y, z)

	// Format 3
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
