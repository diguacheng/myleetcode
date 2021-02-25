package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := 0
	n := len(s)
	i := 0
	var a int
	for i < n {
		a = m[s[i]]
		if i+1 < n && m[s[i+1]] > a {
			res += m[s[i+1]] - a
			i++
		} else {
			res += a
		}
		i++
	}
	return res

}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))

}
