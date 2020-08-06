package main

import (
	"fmt"
	"strings"
)

func palindromePairs(words []string) [][]int {

	res := make([][]int, 0)
	n := len(words)
	for i := 0; i < n; i++ {
		word1 := words[i]
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			word2 := words[j]
			word := word1 + word2
			if isPalindrome(word) {
				res = append(res, []int{i, j})
			}
		}
	}
	return res

}

func isPalindrome(word string) bool {
	s := 0
	e := len(word) - 1
	for s <= e {
		if word[s] != word[e] {
			return false
		}
		s++
		e--
	}
	return true
}

func main() {

	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(words))

}
