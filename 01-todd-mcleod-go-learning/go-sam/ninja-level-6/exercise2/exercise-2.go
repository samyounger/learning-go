package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	f := foo(nums...)
	b := bar(nums)

	fmt.Println(f)
	fmt.Println(b)
}

func foo(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func bar(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
