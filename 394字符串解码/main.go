package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {

	res := ""
	nums := make([]int, 0)
	stack := make([]string, 0)
	i := 0
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			j := i
			for i = j + 1; s[i] >= '0' && s[i] <= '9' && i < len(s); i++ {
			}
			n, _ := strconv.Atoi(s[j:i])
			nums = append(nums, n)
			continue

		} else if s[i] != ']' {
			stack = append(stack, string(s[i]))
		} else {
			temp := ""
			for {
				x := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if x == "[" {
					break
				}
				temp = string(x) + temp
			}
			n := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			temp1 := ""
			for n > 0 {
				temp1 = temp + temp1
				n--
			}
			stack = append(stack, temp1)

		}
		i++
	}
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func main() {
	s := "1000[leetcode]"
	fmt.Println(decodeString(s))
}
