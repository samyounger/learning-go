package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(nums)
	totalSum := sum(nums...)

	fmt.Println(totalSum)

	oddTotal := oddNums(sum, nums...)

	fmt.Println(oddTotal)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func oddNums(f func(x ...int) int, ix ...int) int {
	var evenTotal []int

	for _, v := range ix {
		if v%2 != 0 {
			evenTotal = append(evenTotal, v)
		}
	}
	total := f(evenTotal...)
	return total
}
