package main

import (
	"fmt"
)

func fib(n int) int {
	temp := []int{0, 1}[:]
	/*
		if n==0{
			return 0
		}
		if n==1{
			return 1
		}
		经过分析，发现 上述代码 可以由最后的return 包含住
	*/
	for i := 2; i <= n; i++ {
		temp = append(temp, (temp[i-1]+temp[i-2])%1000000007)
	}
	return temp[n]

}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	res := make([]int, n)
	res[n-1] = 1
	res[n-2] = 2
	for i := n - 3; i >= 0; i-- {
		res[i] = (res[i+1] + res[i+2]) % 1000000007
	}
	return res[0]
}

func main() {
	fmt.Println(numWays(7))

}
