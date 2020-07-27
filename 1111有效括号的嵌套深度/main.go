package main

import "fmt"

func maxDepthAfterSplit(seq string) []int {
	ans := make([]int, 0)
	d := 0
	for _, c := range seq {
		if c == '(' {
			d++
			ans = append(ans, d%2)
		}
		if c == ')' {
			ans = append(ans, d%2)
			d--
		}

	}
	return ans

}
func main() {
	seq := "()(())()"
	fmt.Println(maxDepthAfterSplit(seq))
}
