package main

import (
	"fmt"
)

func permutation(s string) []string {
	res := make(map[string]int)
	temp := make([]byte, 0)
	help([]byte(s), temp, &res)
	ans := make([]string, 0)
	for k := range res {
		ans = append(ans, k)

	}
	return ans

}

func help(s, temp []byte, res *map[string]int) {
	n := len(s)
	if n == 0 {

		(*res)[string(temp)] = 1
	}
	for i := 0; i < n; i++ {

		temp = append(temp, s[i])
		temps := append([]byte{}, s[:i]...)
		temps = append(temps, s[i+1:]...)
		help(temps, temp, res)
		temp = temp[:len(temp)-1]

	}

}
func main() {
	fmt.Println(permutation("aab"))

}
