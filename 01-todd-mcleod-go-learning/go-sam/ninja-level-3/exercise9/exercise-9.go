package main

import "fmt"

func main() {
	c := []string{"foo", "bar", "foobar"}

	for _, s := range c {
		switch s {
		case "foo":
			fmt.Printf("This is 'Foo' => %v\n", s)
		case "bar":
			fmt.Printf("This is 'Bar' => %v\n", s)
		default:
			fmt.Printf("This is default => %v\n", s)
		}
	}
}
