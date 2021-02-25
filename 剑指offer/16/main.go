package main

import "fmt"

func myPow0(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	res := 1.0
	for n > 0 {
		res = res * x
		n--
	}
	return res

}
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n == 0 {
		return 1
	}
	res := 1.0
	if n&1 == 0 {
		// n是偶数
		res = myPow(x, n>>1)
		res *= res
	} else {
		//
		res = myPow(x, (n-1)>>1)
		res = res * res * x
	}

	return res

}

func main() {
	fmt.Printf("%f", myPow(0.00001, 2147483647))

}
