package main

import "fmt"

func longestValidParentheses(s string) int {
	min := 0
	n := len(s)
	stack := make([]byte, 0)
	temp := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		}
		if s[i] == ')' {
			if len(stack) > 0 {
				temp += 2
				stack = stack[:len(stack)-1]
			} else {
				if min < temp {
					min = temp
				}
				temp = 0
			}
		}
	}
	if min < temp {
		min = temp
	}
	return min

}

func main() {
	s1 := "(()"
	s2 := ")()())"
	fmt.Println(longestValidParentheses(s1))
	fmt.Println(longestValidParentheses(s2))

}
