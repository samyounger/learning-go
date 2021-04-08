package main

import "fmt"

func main() {
	y := 1950

	for {
		if y > 2020 {
			break
		}
		fmt.Println(y)
		y++
	}
}
