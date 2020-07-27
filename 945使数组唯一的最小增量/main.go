package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(A []int) int {
	// 超时
	count := 0
	mmap := make(map[int]int)
	for _, v := range A {
		mmap[v]++
	}
	curr := 0
	for len(A) != len(mmap) {
		for i, v := range mmap {
			if v > 1 {
				mmap[i]--
				curr = i + 1
				for {
					count++
					if mmap[curr] == 0 {
						mmap[curr]++
						break
					}
					curr++

				}
			}
		}
	}
	return count
}

func minIncrementForUnique1(A []int) int {
	l := make([]int, 8000)
	sum := 0
	for _, v := range A {
		sum = +v
		l[v]++
	}
	ans := 0
	temp := 0
	for i := 0; i < sum+2; i++ {
		if l[i] > 1 {
			temp = temp + l[i] - 1
			ans = ans - i*(l[i]-1)
		} else if temp > 0 && l[i] == 0 {
			temp--
			ans = ans + i

		}
	}
	return ans
}

func minIncrementForUnique2(A []int) int {
	sort.Ints(A)
	ans := 0
	taken := 0 // 重复的个数
	for i := 1; i < len(A); i++ {
		if A[i-1] == A[i] { // 将操作次数减去A[i],并将重复个数加1
			taken++
			ans = ans - A[i]
		} else {
			// if 不成立，则 区间A[i-1]+1---A[i]-1 都是空的，可以填充数字
			give := min(taken, A[i]-A[i-1]-1)         // 最多可以改变的个数为这两个数的最小值
			ans = ans + give*(give+1)/2 + give*A[i-1] // 相当于一个梯形面积，give*A[i-1]，将give个数加到A[i-1],give*(give+1)/2，相当于，梯形的三角部分
			taken = taken - give

		}
	}
	if taken > 0 {
		// 如果taken 还有剩余，在区间A[len(a)-1]+1-----正无穷之间
		ans = ans + taken*(taken+1)/2 + taken*A[len(A)-1]

	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	A := []int{3, 2, 1, 2, 1, 7}
	fmt.Println(minIncrementForUnique1(A))

}
