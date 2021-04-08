package main

import (
	"fmt"
	"sync"
)

func doSomething(m string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(m)
}

func main() {
	fmt.Println("START")

	var wg sync.WaitGroup

	wg.Add(1)
	go doSomething("HELLO FIRST 1", &wg)
	wg.Add(1)
	go doSomething("HELLO FIRST 2", &wg)

	fmt.Println("HELLO SECOND")

	fmt.Println("HELLO THIRD")

	wg.Wait()
}
