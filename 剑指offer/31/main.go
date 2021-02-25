package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	// 以出栈遍历
	if len(pushed) <= 1 {
		return true
	}
	stack := make([]int, 0)
	i := 0
	p := 0
	stack = append(stack, pushed[i])
	lpop := len(popped) - 1
	lpus := len(pushed) - 1
	for p <= lpop {
		for len(stack) == 0 || popped[p] != stack[len(stack)-1] {
			i++
			if i > lpus {
				return false
			}
			stack = append(stack, pushed[i])
		}
		stack = stack[:len(stack)-1]
		p++

	}
	return true
}

func validateStackSequences1(pushed []int, popped []int) bool {
	// 以入栈遍历
	n := len(pushed)
	stack := make([]int, 0)
	j := 0
	for _, x := range pushed {
		fmt.Println(x)
		stack = append(stack, x)
		for len(stack) != 0 && j < n && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == n
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences1(pushed, popped))

}
