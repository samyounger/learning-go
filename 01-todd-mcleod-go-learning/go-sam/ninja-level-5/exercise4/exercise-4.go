package main

import "fmt"

func main() {
	s := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Foo",
		lastName:  "Bar",
		age:       30,
	}

	fmt.Println(s)
}
