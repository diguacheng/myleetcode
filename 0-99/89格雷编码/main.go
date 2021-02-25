package main

import (
	"fmt"
)

// 镜像反射
// 0 0
// 1 0-1
// 2  00-01-11-10  0(0-1)--1(1-0) 镜像反转
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	list := grayCode(n - 1)
	nl := len(list)
	add := 1 << (n - 1)
	fmt.Println(add)
	for i := nl - 1; i >= 0; i-- {
		x := list[i]
		list = append(list, x|add)
	}
	return list

}

func main() {
	fmt.Println(grayCode(8))

}
