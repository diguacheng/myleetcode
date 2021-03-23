package main

import "fmt"

func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	n := len(input)
	if k > n {
		return nil
	}
	qMink(input, k)
	return input[:k]

}

func qMink(input []int, k int) {
	p := qSelect(input)
	if k > p+1 {
		qMink(input[p+1:], k-p-1)
	} else if k < p {
		qMink(input[:p], k)
	}
	return

}

func qSelect(input []int) int {
	n := len(input)
	left, right := 0, n-1
	p := n - 1
	for left < right {
		for left < right && input[left] <= input[right] {
			left++
		}
		if left < right {
			input[left], input[right] = input[right], input[left]
			p = left
		}
		for left < right && input[left] <= input[right] {
			right--
		}
		if left < right {
			input[left], input[right] = input[right], input[left]
			p = right
		}
	}
	return p
}

func main() {
	// var n int
	// fmt.Scan(&n)
	// arr:=make([]int,n)
	// for i:=0;i<n;i++{
	// 	fmt.Scan(&arr[i])
	// }
	// var k int
	// fmt.Scan(&k)
	var s string
	for {
		fmt.Scan(&s)
		fmt.Println(s)
	}

	//fmt.Println(GetLeastNumbers_Solution(arr,k))
}
