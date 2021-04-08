package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	x := person{
		firstName: "Sam",
		lastName:  "Foo",
		favoriteIceCreamFlavors: []string{
			"vanilla",
			"chocolate",
		},
	}
	y := person{
		firstName: "Cora",
		lastName:  "Bar",
		favoriteIceCreamFlavors: []string{
			"Strawberry",
			"Chocolate",
		},
	}

	fmt.Println(x.firstName, x.lastName)

	for _, f := range x.favoriteIceCreamFlavors {
		fmt.Println(f)
	}

	fmt.Println(y.firstName, y.lastName)

	for _, f := range y.favoriteIceCreamFlavors {
		fmt.Println(f)
	}
}
