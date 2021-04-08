package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	x := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	y := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(x, y)
	fmt.Println(x.doors)
	fmt.Println(y.doors)
}
