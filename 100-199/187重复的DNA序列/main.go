package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n <= 10 {
		return []string{}
	}
	hash := map[string]int{}
	res := []string{}
	var subStr string
	for i := 0; i+10 <= n; i++ {
		subStr = s[i : i+10]
		if c, ok := hash[subStr]; ok && c == 1 {
			res = append(res, subStr)
		}
		hash[subStr]++

	}
	return res

}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))

}