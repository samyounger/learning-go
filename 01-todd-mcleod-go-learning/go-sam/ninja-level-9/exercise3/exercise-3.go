package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doSomething(c *int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("FIRST INCREMENTER VAL: %v\n", *c)
	nc := *c
	runtime.Gosched()
	nc++

	fmt.Printf("NEW COUNTER VAL: %v\t\n", nc)

	*c = nc
}

// FYI - THIS IS A RACE CONDITION
// - NOTE HOW THE FINAL COUNTER DOES NOT EQUAL 10
// - THE FIX USING A MUTEX IS IN EXERCISE 4
// - THE FIX USING ATOMIC IS IN EXERCISE 5
func main() {
	var counter int

	var wg sync.WaitGroup
	gs := 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go doSomething(&counter, &wg)
		// go func() {
		// 	nc := counter
		// 	runtime.Gosched()
		// 	nc++
		// 	counter = nc
		// 	wg.Done()
		// }()
	}

	wg.Wait()

	fmt.Printf("FINAL COUNTER VAL: %v", counter)
}
