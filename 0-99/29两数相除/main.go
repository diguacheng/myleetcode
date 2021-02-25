package main

import "fmt"

func divide(dividend int, divisor int) int {
	flag := true
	if divisor*dividend < 0 {
		flag = false
	}
	if dividend == (-1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := 0
	for dividend >= divisor {
		dividend = dividend - divisor
		res++

	}
	if flag {
		return res
	}
	return -res

}
func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))

}
