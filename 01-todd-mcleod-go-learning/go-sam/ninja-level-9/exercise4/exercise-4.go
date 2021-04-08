package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doSomething(c *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	defer wg.Done()
	fmt.Printf("FIRST INCREMENTER VAL: %v\n", *c)
	nc := *c
	runtime.Gosched()
	nc++

	fmt.Printf("NEW COUNTER VAL: %v\t\n", nc)

	*c = nc
	mu.Unlock()
}

func main() {
	var counter int

	var wg sync.WaitGroup
	var mu sync.Mutex
	gs := 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go doSomething(&counter, &wg, &mu)
		// go func() {
		// 	mu.Lock()
		// 	nc := counter
		// 	runtime.Gosched()
		// 	nc++
		// 	counter = nc
		// 	mu.Unlock()
		// 	wg.Done()
		// }()
	}

	wg.Wait()

	fmt.Printf("FINAL COUNTER VAL: %v", counter)
}
