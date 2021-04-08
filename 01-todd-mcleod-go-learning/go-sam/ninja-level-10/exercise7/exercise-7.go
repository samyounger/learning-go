package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for ix := 0; ix < 10; ix++ {
				c <- ix
			}
			close(c)
		}()
	}

	for v := range c {
		fmt.Println(v)
	}
}
