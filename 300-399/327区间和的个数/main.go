package main

import "fmt"

func countRangeSum(nums []int, lower int, upper int) int {
	// 暴力法
	n := len(nums)
	var sum int
	res := 0
	for i := 0; i < n; i++ {
		sum = 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				res++
			}
		}
	}
	return res
}

func countRangeSum1(nums []int, lower, upper int) int {
	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1 {
			return 0
		}

		n1 := append([]int(nil), arr[:n/2]...)
		n2 := append([]int(nil), arr[n/2:]...)
		cnt := mergeCount(n1) + mergeCount(n2) // 递归完毕后，n1 和 n2 均为有序

		// 统计下标对的数量
		l, r := 0, 0
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		// n1 和 n2 归并填入 arr
		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}
		return cnt
	}

	prefixSum := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}
	return mergeCount(prefixSum)
}

func main() {
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	mergesort(arr)
	fmt.Println(arr)

}
