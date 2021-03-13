package main

import (
	"fmt"
)

func new21Game(N int, K int, W int) float64 {
	list := make(map[int]float64)
	var dp func(cur int) float64
	dp = func(cur int) float64 {
		if ans, ok := list[cur]; ok {
			return ans
		}
		if cur >= K {
			if cur <= N {
				list[cur] = 1.0
				return 1.0
			}
			list[cur] = 0.0
			return 0.0
		}
		ans := 0.0
		for i := W; i >= 1; i-- {
			ans = ans + dp(cur+i)
		}
		ans = (1.0 / float64(W)) * ans
		list[cur] = ans
		return ans
	}
	defer fmt.Println(list)
	return dp(0)

}
func new21Game1(N int, K int, W int) float64 {
	list := make([]float64, K+1)
	var dp func(cur int) float64
	dp = func(cur int) float64 {
		if cur >= K {
			if cur <= N {
				return 1.0
			}
			return 0.0
		}
		ans := list[cur]
		if ans > 0.0 {
			return ans
		}
		for i := W; i >= 1; i-- {
			ans = ans + dp(cur+i)
		}
		ans = (1.0 / float64(W)) * ans
		list[cur] = ans
		return ans
	}
	return dp(0)

}

func main() {
	fmt.Println(new21Game(21, 17, 10))
	fmt.Println(new21Game1(21, 17, 10))

}
