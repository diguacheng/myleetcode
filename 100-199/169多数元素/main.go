package main

import (
	"math/rand"
	"sort"
)

func majorityElement(nums []int) int {
	// 哈希法
	res := make(map[int]int, 10)
	l := len(nums) / 2
	for _, num := range nums {
		res[num] = res[num] + 1
		if res[num] > l {
			return num
		}
	}
	return 0
}

func majorityElement1(nums []int) int {
	// 对数组排序，返回第[n/2]个数
	sort.Ints(nums)
	return nums[len(nums)/2]

}

func majorityElement2(nums []int) int {
	// 随机取一个数，检查它是不是符合条件
	for {
		x := nums[rand.Intn(len(nums))]
		count := 0
		for _, num := range nums {
			if num == x {
				count++
			}
		}
		if count > len(nums)/2 {
			return x
		}
	}

}

func majorityElement3(nums []int) int {
	// Boyer-Moore 算法
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--

		}
	}
	return candidate

}

func main() {

}
