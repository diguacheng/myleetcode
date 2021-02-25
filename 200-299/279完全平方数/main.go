package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	// 100 ms
	if n <= 3 {
		return n
	}
	dp := make([]int, n+1)
	list := []int{}
	var x int
	var i int
	for i = 1; i <= n; i++ {
		x = i * i
		if x > n {
			break
		}
		dp[x] = 1
		list = append(list, x)
	}
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 1
	lenl := len(list)
	var help func(n int) int
	help = func(n int) int {
		if dp[n] != 0 {
			return dp[n]
		}
		min := n + 1
		for i := lenl - 1; i >= 0; i-- {
			if n-list[i] > 0 {
				if min > help(n-list[i]) {
					min = help(n - list[i])
				}
			}
		}
		dp[n] = min + 1
		return dp[n]
	}
	return help(n)
}

func numSquares1(n int) int {
	// 1216ms
	if n <= 3 {
		return n
	}
	table := map[int]int{}
	var x int
	var i int
	for i = 1; i <= n; i++ {
		x = i * i
		if x > n {
			break
		}
		table[x] = 1
	}
	i = i - 2
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		if table[i] == 1 {
			dp[i] = 1
			continue
		}
		var min = n + 1
		for j := 1; j < i; j++ {
			if dp[j]+dp[i-j] < min {
				min = dp[j] + dp[i-j]
			}
		}
		dp[i] = min

	}
	return dp[n]
}

func numSquares2(n int) int {
	//12 ms
	if n <= 3 {
		return n
	}
	table := map[int]int{}
	var x int
	var i int
	for i = 1; i <= n; i++ {
		x = i * i
		if x > n {
			break
		}
		table[x] = 1
	}

	var isDividedBy func(n, count int) bool
	isDividedBy = func(n, count int) bool {
		if count == 1 {
			return table[n] > 0
		}
		for k := range table {
			if isDividedBy(n-k, count-1) {
				return true
			}
		}
		return false
	}
	for i := 1; i <= n; i++ {
		if isDividedBy(n, i) {
			return i
		}
	}
	return 0
}

func numSquares3(n int) int {
	//60 ms
	if n <= 3 {
		return n
	}
	list := make([]int, 0)
	var x int
	var i int
	for i = 1; i <= n; i++ {
		x = i * i
		if x > n {
			break
		}
		list = append(list, x)
	}
	lenl := len(list)
	level := 0
	set := make(map[int]int)
	var nextSet map[int]int
	set[n] = 1
	var sqrNum int
	for len(set) > 0 {
		level++
		nextSet = map[int]int{}
		for k := range set {
			for i := 0; i < lenl; i++ {
				sqrNum = list[i]
				if k == sqrNum {
					return level
				}
				if k > sqrNum {
					nextSet[k-sqrNum] = 1
				} else {
					break
				}
			}

		}
		set = nextSet
	}
	return level

}

func numSquares4(n int) int {

	isSquare := func(c int) bool {
		sq := int(math.Sqrt(float64(c)))
		return sq*sq == c
	}

	for n&3 == 0 {
		n = n >> 2
	}
	if n&7 == 7 {
		return 4
	}
	if isSquare(n) {
		return 1
	}
	s := int(math.Sqrt(float64(n))) + 1
	for i := 0; i < s; i++ {
		if isSquare(n - i*i) {
			return 2
		}

	}
	return 3

}

func main() {
	fmt.Println(numSquares3(12))
	fmt.Println(12 & 3)
	fmt.Println(numSquares4(12))

}
