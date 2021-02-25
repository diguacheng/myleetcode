package main

import "fmt"

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	res := min(height[start], height[end]) * (end - start)
	for start < end {
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
		temp := min(height[start], height[end]) * (end - start)
		if res < temp {
			res = temp
		}

	}
	return res

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

}
