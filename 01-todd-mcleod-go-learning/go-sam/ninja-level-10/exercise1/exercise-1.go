package main

import (
	"fmt"
)

func main() {
	// Using func literal, aka, anonymous self-executing func
	// c := make(chan int, 1)
	c := make(chan int)

	// Using func literal, aka, anonymous self-executing func
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
