package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := make([]string, 0)
	end := len(s) - 1
	start := end
	for end >= 0 {
		if s[end] == ' ' {
			end--
		} else {
			start = end
			for start >= 0 && s[start] != ' ' {
				start--
			}
			res = append(res, s[start+1:end+1])
			end = start
		}
	}
	ss := ""
	for i, v := range res {
		ss = ss + v
		if i != len(res)-1 {
			ss = ss + " "

		}
	}
	return ss
}

func reverseWords1(s string) string {
	seg := strings.Fields(s)
	resseg := make([]string, 0)
	for i := len(seg) - 1; i >= 0; i-- {
		resseg = append(resseg, seg[i])
	}
	return strings.Join(resseg, " ")

}
func main() {
	s := "a good   example "
	fmt.Println(reverseWords1(s))

}
