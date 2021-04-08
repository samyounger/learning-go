package main

import "fmt"

func main() {
	x := []string{"The Neuk", "The Aeolian", "Whitslaid"}

	for _, s := range x {
		if s == "The Aeolian" {
			fmt.Printf("We're home at %v\n", s)
		} else {
			fmt.Printf("We're somewhere else, at %v\n", s)
		}
	}
}
