package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
	width  float64
}

type shape interface {
	area() float64
}

func main() {
	s := square{
		length: 4,
		width:  5,
	}
	c := circle{
		radius: 15,
	}

	info(s)
	info(c)
}

func (c circle) area() float64 {
	p := 3.14159265359
	return math.Pow((p * c.radius), 2)
}

func (s square) area() float64 {
	return s.length * s.width
}

func info(s shape) {
	a := s.area()
	fmt.Println(a)
}
