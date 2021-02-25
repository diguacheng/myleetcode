package main

import "fmt"

func lengthOfLastWord(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	end := n - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end == -1 {
		return 0
	}
	start := end - 1
	for start >= 0 {
		if s[start] == ' ' {
			return end - start
		}
		start--
	}
	return end + 1

}

func main() {
	fmt.Println(lengthOfLastWord("a"))

}
