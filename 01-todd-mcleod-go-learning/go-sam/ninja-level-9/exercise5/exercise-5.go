package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func doSomething(c *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(c, 1)
	atomic.LoadInt64(c)
}

func main() {
	var counter int64

	var wg sync.WaitGroup
	gs := 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go doSomething(&counter, &wg)
		// go func() {
		// 	atomic.AddInt64(&counter, 1)
		// 	atomic.LoadInt64(&counter)
		// 	wg.Done()
		// }()
	}

	wg.Wait()

	fmt.Printf("FINAL COUNTER VAL: %v", counter)
}
