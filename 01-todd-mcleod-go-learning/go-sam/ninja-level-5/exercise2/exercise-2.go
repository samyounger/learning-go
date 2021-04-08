package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "Sam",
		lastName:  "Foo",
		favoriteIceCreamFlavors: []string{
			"vanilla",
			"chocolate",
		},
	}
	p2 := person{
		firstName: "Cora",
		lastName:  "Bar",
		favoriteIceCreamFlavors: []string{
			"Strawberry",
			"Chocolate",
		},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for _, val := range v.favoriteIceCreamFlavors {
			fmt.Println(val)
		}
	}
}
