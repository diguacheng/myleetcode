package main

import (
	"fmt"
)

var res []float64

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {

	x1, y1 := start1[0], start1[1]
	x2, y2 := end1[0], end1[1]
	x3, y3 := start2[0], start2[1]
	x4, y4 := end2[0], end2[1]
	res = make([]float64, 0)
	// 判断是否平行
	if (y4-y3)*(x2-x1) == (y2-y1)*(x4-x3) {
		// 若平行 判断 （x3，y3)是否在直线（x1，y1）（x2，y2）上
		if (y2-y1)*(x3-x1) == (y3-y1)*(x2-x1) {
			// 判断（x3，y3）是否在线段（x1,y1)--(x2,y2)上
			if inside(x1, y1, x2, y2, x3, y3) {
				update(&res, float64(x3), float64(y3))
			}

			// 判断（x4，y4）是否在线段（x1,y1)--(x2,y2)上
			if inside(x1, y1, x2, y2, x4, y4) {
				update(&res, float64(x4), float64(y4))
			}

			// 判断（x1，y1）是否在线段（x3,y3)--(x4,y4)上
			if inside(x3, y3, x4, y4, x1, y1) {
				update(&res, float64(x1), float64(y1))
			}

			// 判断（x2，y2）是否在线段（x3,y3)--(x4,y4)上
			if inside(x3, y3, x4, y4, x2, y2) {
				update(&res, float64(x2), float64(y2))
			}
		}
	} else {
		t1 := float64((x3*(y4-y3) + y1*(x4-x3) - y3*(x4-x3) - x1*(y4-y3))) / float64(((x2-x1)*(y4-y3) - (x4-x3)*(y2-y1)))
		t2 := float64((x1*(y2-y1) + y3*(x2-x1) - y1*(x2-x1) - x3*(y2-y1))) / float64(((x4-x3)*(y2-y1) - (x2-x1)*(y4-y3)))
		if t1 >= 0.0 && t1 <= 1.0 && t2 >= 0.0 && t2 <= 1.0 {
			res = []float64{float64(x1) + t1*float64((x2-x1)), float64(y1) + t1*float64((y2-y1))}
		}
	}
	return res

}

func update(res *[]float64, xk, yk float64) {

	if len(*res) == 0 || xk < (*res)[0] || (xk == (*res)[0] && yk < (*res)[1]) {
		*res = []float64{xk, yk}
	}
}

//  判断(xk,yk)是否在线段（x1，y1)--(x2,y2)上 前提是(xk,yk)在直线（x1，y1)--(x2,y2)上
func inside(x1, y1, x2, y2, xk, yk int) bool {
	// 与x轴平行 判断x部分
	// 与y轴平行 判断y部分
	// 反之 都要判断
	return (x1 == x2 || (min(x1, x2) <= xk && xk <= max(x1, x2))) && (y1 == y2 || (min(y1, y2) <= yk && yk <= max(y1, y2)))

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
	start1 := []int{0, 0}
	start2 := []int{1, 0}
	end1 := []int{1, 1}
	end2 := []int{2, 1}
	fmt.Println(intersection(start1, end1, start2, end2))

}
