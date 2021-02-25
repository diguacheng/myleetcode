package main

import (
	"fmt"
)

// func minArray(numbers []int) int {
// 	start,end := 0, len(numbers) - 1
// 	for start<end{
// 		mid:=(start+end)/2
// 		if numbers[mid]>numbers[end]{
// 			start=mid+1
// 		}else if numbers[mid]<numbers[start]{
// 			end=mid
// 		}else{
// 			end--
// 		}
// 	}
// 	return numbers[start]
// }

func minArray(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	left, right := 0, n-1
	for left < right {
		mid := left + (right-right)/2
		if numbers[mid] >= numbers[n-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return numbers[left]

}

func main() {
	numbers := []int{1, 3, 5}
	fmt.Println(minArray(numbers))

}
