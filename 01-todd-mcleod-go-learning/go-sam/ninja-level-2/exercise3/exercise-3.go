package main

import "fmt"

const unTypedConst = 42
const typedConst int = 42

// Alternative
// const (
// 	a = 42
// 	b int = 43
// )

func main() {
	fmt.Println(unTypedConst)
	fmt.Println(typedConst)
}
