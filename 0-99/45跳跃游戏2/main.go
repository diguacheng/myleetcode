package main

import (
	"fmt"
)

func jump(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps

}

func jump1(nums []int) int {
	ans := 0
	start := 0
	end := 1
	for end < len(nums) {
		maxpos := 0
		for i := start; i < end; i++ {
			// 优化：i是从头跑到尾的
			fmt.Println(i)
			maxpos = max(maxpos, i+nums[i])

		}
		start = end
		end = maxpos + 1
		ans++
	}
	return ans

}

func jump2(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {

		maxPosition = max(maxPosition, i+nums[i])
		if i == end {

			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump1(nums))

}
