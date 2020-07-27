package main

import "fmt"

func climbStairs(n int) int {
	res := make([]int, n+1)
	res[n] = 1
	res[n-1] = 1
	n = n - 2
	for n >= 0 {
		res[n] = res[n+1] + res[n+2]
		n--
	}
	return res[0]

}

func main() {
	fmt.Println(climbStairs(2))

}
