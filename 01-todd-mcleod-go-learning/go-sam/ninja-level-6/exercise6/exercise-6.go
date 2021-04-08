package main

import "fmt"

func main() {
	func() {
		fmt.Println("Here in anonymous func")
	}()
}
