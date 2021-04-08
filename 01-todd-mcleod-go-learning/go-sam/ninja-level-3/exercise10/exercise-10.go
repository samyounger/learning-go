package main

import "fmt"

func main() {
	fmt.Printf("true, %v\n", true && true)
	fmt.Printf("false, %v\n", true && false)
	fmt.Printf("true, %v\n", true || true)
	fmt.Printf("true, %v\n", true || false)
	fmt.Printf("false, %v\n", !true)
}
