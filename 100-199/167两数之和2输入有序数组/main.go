package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	var temp int
	for i := 0; i < n-1; i++ {
		temp = target - numbers[i]
		for j := i + 1; j < n; j++ {
			if numbers[j] == temp {
				return []int{i + 1, j + 1}
			}

		}

	}
	return []int{0, 0}

}

func twoSum1(numbers []int, target int) []int {
	n := len(numbers)
	hash := map[int]int{}
	var temp int
	for i := 0; i < n; i++ {
		temp = target - numbers[i]
		if j, ok := hash[temp];   ok {
			return []int{ j + 1,i + 1}
		}
		hash[numbers[i]] = i

	}
	return []int{0, 0}

}

func main() {
	s := []int{2, 7, 11, 15}
	fmt.Println(twoSum1(s,9))

}