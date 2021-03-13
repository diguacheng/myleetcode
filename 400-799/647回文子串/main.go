package main

import "fmt"

func countSubstrings3(s string) int {
	// 中心拓展法 速度最快
	res := 0

	n := len(s)
	var help func(start, end int) int

	help = func(start, end int) int {
		sum := 0
		for left, right := start, end; left >= 0 && right < n && s[left] == s[right]; left, right = left-1, right+1 {
			sum++
		}
		return sum
	}
	for i := 0; i < n; i++ {
		res += help(i, i)
		res += help(i, i+1)
	}
	return res
}

func countSubstrings(s string) int {

	table := map[string]int{}
	check := func(s string) bool {
		n := len(s)
		if n == 1 {
			return true
		}
		if table[s] == 1 {
			return true
		}
		if table[s] == -1 {
			return false
		}

		l := n / 2
		for i := 0; i < l; i++ {
			if s[i] != s[n-1-i] {
				table[s] = -1
				return false
			}
		}
		table[s] = 1
		return true
	}
	n := len(s)
	res := n
	isVisited := make([][]int, n+1)
	for i := 0; i < n; i++ {
		isVisited[i] = make([]int, n+1)
	}
	var backTrack func(i int)
	backTrack = func(i int) {
		if i+2 > n {
			return
		}
		for j := i + 2; j <= n; j++ {
			if check(s[i:j]) && isVisited[i][j] == 0 {
				isVisited[i][j] = 1
				res++
				backTrack(j)
			}

		}
	}
	backTrack(0)
	return res
}

func countSubstrings1(s string) int {

	table := map[string]int{}
	check := func(s string) bool {
		n := len(s)
		if n == 1 {
			return true
		}
		if table[s] == 1 {
			return true
		}
		if table[s] == -1 {
			return false
		}

		l := n / 2
		for i := 0; i < l; i++ {
			if s[i] != s[n-1-i] {
				table[s] = -1
				return false
			}
		}
		table[s] = 1
		return true
	}
	n := len(s)
	res := n
	gap := 2
	var end int
	for gap <= n {
		end = n - gap
		for i := 0; i <= end; i++ {
			if check(s[i : i+gap]) {
				res++
			}
		}
		gap++
	}
	return res
}

func countSubstrings2(s string) int {

	table := map[string]int{}
	check := func(s string) bool {
		n := len(s)
		if n == 1 {
			return true
		}
		if table[s] == 1 {
			return true
		}
		if table[s] == -1 {
			return false
		}

		l := n / 2
		for i := 0; i < l; i++ {
			if s[i] != s[n-1-i] {
				table[s] = -1
				return false
			}
		}
		table[s] = 1
		return true
	}
	n := len(s)
	res := n
	isVisited := make([][]int, n+1)
	for i := 0; i < n; i++ {
		isVisited[i] = make([]int, n+1)
	}
	var backTrack func(i int)
	backTrack = func(i int) {
		if i+2 > n {
			return
		}
		for j := i + 2; j <= n; j++ {
			if check(s[i:j]) && isVisited[i][j] == 0 {
				isVisited[i][j] = 1
				res++
				backTrack(j)
			}

		}
	}
	backTrack(0)
	return res
}

func main() {
	fmt.Println(countSubstrings3("12343435342341"))
}
