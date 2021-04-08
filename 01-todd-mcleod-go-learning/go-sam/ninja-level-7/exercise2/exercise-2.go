package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "ham",
		last:  "salot",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	p.first = "foo"
	// (*p).first = "foo"
}
