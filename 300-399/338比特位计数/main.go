package main

import "fmt"

func countBits(num int) []int {
	if num == 0 {
		return []int{0}

	}
	res := make([]int, num+1)
	res[1] = 1
	a := 1
	index := 2
	for index <= num {
		a = a << 1
		for i := 0; i < a && index <= num; i++ {
			res[index] = res[i] + 1
			index++
		}

	}
	return res

}

func main() {
	fmt.Println(countBits(1))
	fmt.Println(countBits(5))

}
