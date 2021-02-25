package main

import "fmt"

func lastRemaining1(n int, m int) int {
	// 暴力法 超时
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = i
	}
	i := 0
	for len(l) > 1 {
		i = (i + m - 1) % len(l)
		l = append(l[:i], l[i+1:]...)
	}
	return l[0]
}
func lastRemaining(n int, m int) int {
	// 数学+递归
	if n == 1 {
		return 0
	}
	x := lastRemaining(n-1, m)
	return (m + x) % n

}

func lastRemaining2(n int, m int) int {
	// 数学+迭代
	res := 0
	for i := 2; i != n+1; i++ {
		res = (res + m) % i
	}
	return res

}
func main() {
	fmt.Println(lastRemaining2(1, 17))

}
