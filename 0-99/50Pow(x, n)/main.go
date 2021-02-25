package main

import "fmt"

func myPow(x float64, n int) float64 {
	// 快速幂+递归
	ans := 1.0
	flag := true
	if n < 0 {
		n = -n
		flag = false
	}
	if n == 0 {
		return ans
	}
	temp := myPow(x, n/2)
	if n%2 == 1 {
		ans = temp * temp * x

	} else {
		ans = temp * temp
	}
	if flag {
		return ans
	}
	return 1 / ans

}

func myPow1(x float64, n int) float64 {
	// 快速幂+迭代   将n 转换为2进制，从左向右 遍历 ，某位为1 就更新ans
	ans := 1.0
	flag := true
	if n < 0 {
		n = -n
		flag = false
	}
	if n == 0 {
		return ans
	}
	Contribute := x
	for n > 0 {
		if n%2 == 1 {
			ans = ans * Contribute
		}
		Contribute = Contribute * Contribute
		n = n / 2

	}
	if flag {
		return ans
	}
	return 1 / ans

}

func main() {
	fmt.Println(myPow1(2.0, -10))

}
