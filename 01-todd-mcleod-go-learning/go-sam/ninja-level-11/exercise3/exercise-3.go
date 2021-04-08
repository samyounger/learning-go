package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("An error occurred: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	ce := customErr{
		info: "Foobar custom error",
	}

	foo(ce)
}
