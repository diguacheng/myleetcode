package main

import "fmt"

func smallestSubsequence(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++ // left 记录剩余元素的数量，当删除栈中某个元素的时候
		// 如果后面已经没有这个元素了 就不删除。
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		// 判断该元素是否加入到了栈中
		if !inStack[ch-'a'] {
			// 若栈顶元素的大于当前元素 就将栈顶元素出栈 贪心思想
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}

func main() {
	fmt.Println(smallestSubsequence("bcabc"))
}
