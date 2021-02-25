package main

import "fmt"

func validPalindrome(s string) bool {
	flag := 1
	var help func(s string) bool
	help = func(s string) bool {
		l, r := 0, len(s)-1
		for l <= r {
			if s[l] == s[r] {
				l++
				r--
			} else if flag > 0 {
				flag--
				return help(s[l:r]) || help(s[l+1:r+1])
			} else {
				return false
			}

		}
		return true

	}
	return help(s)

}

func main() {
	fmt.Println(validPalindrome("abca"))

}
