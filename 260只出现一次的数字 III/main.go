package main

import "fmt"

func singleNumber(nums []int) []int {
	flag := 0
	for _, num := range nums {
		flag = flag ^ num
	}
	diviser := 1
	for flag&diviser == 0 {
		diviser = diviser << 1
	}
	var a, b int
	for _, num := range nums {
		if diviser&num == 0 {
			a = a ^ num
		} else {
			b = b ^ num
		}
	}
	return []int{a, b}
}

func main() {
	l := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(l))
}