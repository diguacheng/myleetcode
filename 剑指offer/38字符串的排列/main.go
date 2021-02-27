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
		//temps := append([]byte{}, s[:i]...)
		//temps = append(temps, s[i+1:]...)
		temps := string(s[:i]) + string(s[i+1:])
		help([]byte(temps), temp, res)
		temp = temp[:len(temp)-1]

	}

}
func main() {
	fmt.Println(permutation("addab"))

}

func permutation1(s string) []string {
	res := make([]string, 0)
	dfs(res, []byte(s), 0)
	return res
}

func dfs(res []string, s []byte, pos int) {
	n := len(s)
	if pos == n {
		res = append(res, string(s))
	}
	for i := pos; i < n; i++ {
		flag := true
		for j := pos; j < i; j++ {
			if s[j] == s[i] {
				flag = false
			}
		}
		if flag {
			s[pos], s[i] = s[i], s[pos]
			dfs(res, s, pos+1)
			s[pos], s[i] = s[i], s[pos]

		}
	}
}
