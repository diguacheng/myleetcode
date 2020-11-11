package main

import "fmt"

func getMoneyAmount(n int) int {
	var help func(int, int) int
	help = func(start, end int) int {
		if start == end {
			return 0
		}
		mid := (start+end+1)/2
		fmt.Println(mid)
		a,b:=help(start,mid-1),help(mid+1,end)
		return mid+max(a,b)
	}
	return help(1, n)

}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getMoneyAmount(10))
}