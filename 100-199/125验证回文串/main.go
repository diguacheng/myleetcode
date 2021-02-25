package main

import (
	"regexp"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	reg := regexp.MustCompile(`[A-Za-z0-9]+`)
	s = strings.Join(reg.FindAllString(s, -1), "")
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--

	}

	return true

}

func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	var isvalid func(c rune) bool
	isvalid = func(c rune) bool {

		return unicode.IsDigit(c) || unicode.IsLetter(c)
	}

	l, r := 0, len(s)-1
	for l < r {
		lb, rb := isvalid(rune(s[l])), isvalid(rune(s[r]))
		if !lb && !rb {
			l++
			r--
		} else if !lb {
			l++
		} else if !rb {
			r--
		} else if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}

	}
	return true
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")

}
