package main

import "fmt"

func longestPalindrome(s string) int {
	hashmap := make(map[rune]int)
	for _, r := range s {
		hashmap[r]++
	}
	res := 0
	count := 0
	for _, v := range hashmap {
		if v%2 == 0 {
			res = res + v
		} else {
			count = 1
			res = res + v - 1
		}
	}
	if count == 1 {
		res = res + 1
	}
	return res

}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
