package main

import "fmt"

func integerReplacement(n int) int {
	dplist := make(map[int]int)
	var dp func(i int) int
	dp = func(i int) int {
		if i == 1 {
			return 0
		}
		if i >= 2 && dplist[i] != 0 {
			return dplist[i]
		}
		if i%2 == 0 {
			dplist[i] = dp(i/2) + 1
			return dplist[i]
		}
		minl := min(dp(i-1), dp(i+1))
		dplist[i] = minl + 1
		return dplist[i]
	}
	return dp(n)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { 
	fmt.Println(integerReplacement(100000000)) 
	fmt.Println(integerReplacement(7)) 
	fmt.Println(integerReplacement(6)) 
}