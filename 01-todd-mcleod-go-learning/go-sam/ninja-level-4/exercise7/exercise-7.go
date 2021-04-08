package main

import "fmt"

func main() {
	x := [][]string{
		{
			"James",
			"Bond",
			"Shaken, not stirred",
		},
		{
			"Miss",
			"Moneypenny",
			"Helloooooo, James.",
		},
	}

	for i := range x {
		for a := range x[i] {
			fmt.Println(x[i][a])
		}
	}
}
