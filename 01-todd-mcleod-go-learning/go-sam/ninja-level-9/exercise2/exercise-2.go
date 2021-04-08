package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p *person) speak() {
	fmt.Printf("Hello, my name is %v %v\n", p.firstName, p.lastName)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		firstName: "Foo",
		lastName:  "Bar",
	}

	fmt.Println("One")
	saySomething(&p)

	fmt.Println("TWO")
	p.speak()
}
