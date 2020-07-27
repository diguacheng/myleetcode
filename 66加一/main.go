package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	temp := 1
	for i := n - 1; i >=0; i-- {
		if temp == 1 {
			if digits[i] == 9 {
				digits[i] = 0
				if i==0{
					digits=append([]int{1}, digits...)
					return digits
				}

			} else {
				digits[i]++
				temp--
				
			}
		}
	}
	return digits
}

func main() {
	a := []int{9}
	fmt.Println(plusOne(a))
}
