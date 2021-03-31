package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	count := 0
	start, end := 0, 0
	for end < n {
		for end < n && nums[end] == 0 {
			count++
			end++
		}
		if end < n {
			nums[start] = nums[end]
			start++
			end++
		}

	}

	for start < n {
		nums[start] = 0
		start++
	}

}

func moveZeroes1(nums []int) {
    left, right, n := 0, 0, len(nums)
    for right < n {
		fmt.Println(nums)
        if nums[right] != 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
        right++
    }
}



func main() {
	s := []int{1,2,0,3,4}
	moveZeroes1(s)
	fmt.Println(s)

}
