package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""

	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	minl := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minl > len(strs[i]) {
			minl = len(strs[i])

		}
	}
	var temp byte
	for i < minl {
		temp = strs[0][i]
		for j := 1; j < len(strs); j++ {
			if temp != strs[j][i] {
				return strs[0][:i+1]

			}

		}
		i++

	}
	return strs[0][:i+1]

}
func main() {
	strs := []string{"a", "b"}
	fmt.Println(longestCommonPrefix(strs))
}
