package main

import "fmt"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	res := area(A, B, C, D) + area(E, F, G, H)
	y := min(D, H) - max(B, F)
	x := min(C, G) - max(A, E)
	if x > 0 && y > 0 {
		res = res - x*y
	}
	return res

}

func area(x1, y1, x2, y2 int) int {
	return (y2 - y1) * (x2 - x1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))

}
