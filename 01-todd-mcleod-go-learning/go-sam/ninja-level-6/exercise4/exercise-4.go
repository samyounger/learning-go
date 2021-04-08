package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Foo",
		last:  "Bar",
		age:   100,
	}
	p.speak()
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and I am %v", p.first, p.last, p.age)
}
