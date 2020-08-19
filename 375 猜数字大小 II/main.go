package main

import "fmt"

func getMoneyAmount(n int) int {

	var help func(int, int) int
	help = func(start, end int) int {
		if start +1== end {
			return end
		}
		mid := (start + end)/2
		return max(help(start, mid-1), help(mid+1, end)) + mid
	}
	return help(1, n)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getMoneyAmount(10))
}