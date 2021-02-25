package main

import (
	"container/heap"
	"fmt"
)

func nthUglyNumber(n int) int {
	// 暴力法  效率太慢
	list := make([]int, 0)
	i := 1
	for len(list) < n {
		if isUgly(i) {
			list = append(list, i)
		}
		i++

	}
	return list[len(list)-1]

}

func isUgly(n int) bool {
	for n%2 == 0 {
		n = n / 2
	}
	for n%3 == 0 {
		n = n / 3
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n == 1
}

func nthUglyNumber1(n int) int {
	// 将计算过的树 存入数组中，利用之前的结果 计算
	if n <= 0 {
		return 0
	}
	list := make([]int, n)
	list[0] = 1
	nextindex := 1
	i, j, k := 0, 0, 0
	for nextindex < n {
		x := min(min(list[i]*2, list[j]*3), list[k]*5)
		list[nextindex] = x
		for list[i]*2 <= list[nextindex] {
			i++
		}
		for list[j]*3 <= list[nextindex] {
			j++
		}
		for list[k]*5 <= list[nextindex] {
			k++
		}
		nextindex++

	}
	return list[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type maxHeap []int

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}

func nthUglyNumber2(n int) int {
	// 堆排序
	if n <= 0 {
		return 0
	}
	seen := make(map[int]int)
	seen[1] = 1
	list := [1690]int{}
	h := &maxHeap{1}
	l := [3]int{2, 3, 5}
	for i := 0; i < n; i++ {
		x := heap.Pop(h).(int)
		for _, v := range l {
			if seen[x*v] == 0 {
				seen[x*v] = 1
				heap.Push(h, x*v)
			}

		}
		list[i] = x

	}
	return list[n-1]
}

func nthUglyNumber3(n int) int {
	// 动态规划
	// 	官方题解里提到的三个指针p2，p3，p5，但是没有说明其含义，实际上pi的含义是有资格同i相乘的最小丑数的位置。这里资格指的是：如果一个丑数nums[pi]通过乘以i可以得到下一个丑数，那么这个丑数nums[pi]就永远失去了同i相乘的资格（没有必要再乘了），我们把pi++让nums[pi]指向下一个丑数即可。

	// 不懂的话举例说明：

	// 一开始，丑数只有{1}，1可以同2，3，5相乘，取最小的1×2=2添加到丑数序列中。

	// 现在丑数中有{1，2}，在上一步中，1已经同2相乘过了，所以今后没必要再比较1×2了，我们说1失去了同2相乘的资格。

	// 现在1有与3，5相乘的资格，2有与2，3，5相乘的资格，但是2×3和2×5是没必要比较的，因为有比它更小的1可以同3，5相乘，所以我们只需要比较1×3，1×5，2×2。

	// 依此类推，每次我们都分别比较有资格同2，3，5相乘的最小丑数，选择最小的那个作为下一个丑数，假设选择到的这个丑数是同i（i=2，3，5）相乘得到的，所以它失去了同i相乘的资格，把对应的pi++，让pi指向下一个丑数即可。

	if n <= 0 {
		return 0
	}
	list := [1690]int{1}
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		num := min(min(list[i2]*2, list[i3]*3), list[i5]*5)
		list[i] = num
		if num == list[i2]*2 {
			i2++
		}
		if num == list[i3]*3 {
			i3++
		}
		if num == list[i5]*5 {
			i5++
		}
	}
	return list[n-1]
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, nthUglyNumber2(i))
	}

}
