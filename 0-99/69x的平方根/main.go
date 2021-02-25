package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	// 袖珍计算器算法
	if x == 0 {
		return 0
	}
	fmt.Println(math.Exp(0.5 * math.Log(float64(x))))
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	fmt.Println(ans)
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans

}

func mySqrt1(x int) int {
	// 二分法
	start, end := 0, x
	ans := 0
	for start <= end {
		mid := (start + end) / 2
		if mid*mid <= x {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return ans
}

func mySqrt2(x int) int {
	// 牛顿法
	if x == 0 {
		return 0
	}
	c, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + c/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
func main() {
	fmt.Println(mySqrt1(8))

}
