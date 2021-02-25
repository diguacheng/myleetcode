package main

import "fmt"

func canMeasureWater(x int, y int, z int) bool {
	// 这道题主要是通过数学定理来证明 mx+ny=z 如果x，y 是相对素数，则可以生成任意的z，反之，mx+ny 只能等于x,y 最大公因数的倍数。
	// 或者说，相对素数的最大公因数是1，任何数都是1的倍数。

	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || z == x+y
	}
	if x < y {
		x, y = y, x
	}
	return z%gcd(x, y) == 0

}
func gcd(x, y int) int {
	if x%y == 0 {
		return y
	}
	return gcd(y, x%y)
}

func main() {
	fmt.Println(canMeasureWater(3, 5, 4))

}
