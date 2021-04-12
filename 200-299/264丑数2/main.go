package main

import "fmt"

func nthUglyNumber(n int) int {
	if n <= 5 {
		return n
	}
	arr := make([]int, n)
	arr[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		x := min(min(2*arr[i2], arr[i3]*3), arr[i5]*5)
		arr[i] = x
		for 2*arr[i2] <= x {
			i2++
		}
		for arr[i3]*3 <= x {
			i3++
		}
		for arr[i5]*5<= x {
			i5++
		}
	}
	return arr[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n:=11
	nthUglyNumber(n)
}