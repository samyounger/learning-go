package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range x {
		fmt.Printf("Last name: %v\n", k)
		for i, f := range v {
			fmt.Printf("\tFavourite thing of index %v:\n\t\t- %v\n", i, f)
		}
	}
}
