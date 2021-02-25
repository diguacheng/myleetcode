package main

import "fmt"

func subarraySum(nums []int, k int) int {
	// 暴力法
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count

}

func subarraySum1(nums []int, k int) int {
	// 前缀和
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	count := 0
	for left := 0; left < n; left++ {
		for right := left; right < n; right++ {
			if preSum[right+1]-preSum[left] == k {
				count++
			}
		}
	}
	return count
}
func subarraySum2(nums []int, k int) int {
	// 前缀和、哈希优化
	n := len(nums)
	preSum := make([]int, n+1)
	m := make(map[int]int)
	m[0] = 1
	preSum[0] = 0
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
		m[preSum[i+1]]++ // key为前缀和，value为前缀和的个数
	}
	count := 0
	for i := 0; i < n; i++ {
		if _, ok := m[preSum[i+1]-k]; ok {
			// preSum[i+1]表示前i个元素的和，满足pre[i+1]-pre[j]=k，则
			// [j,i]的和为k。即 p[j]=pre[i+1]-k,，只要找到了pre[i+1]-k的个数，就知道以i为尾，满足条件的个数
			// 因此 可以用map进行优化。
			count += m[preSum[i+1]-k]
		}
	}
	return count
}

func subarraySum3(nums []int, k int) int {
	// 前缀和、哈希优化 进一步优化
	n := len(nums)
	m := make(map[int]int)
	m[0] = 1
	pre, count := 0, 0
	for i := 0; i < n; i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			// preSum[i+1]表示前i个元素的和，满足pre[i+1]-pre[j]=k，则
			// [j,i]的和为k。即 p[j]=pre[i+1]-k,，只要找到了pre[i+1]-k的个数，就知道以i为尾，满足条件的个数
			// 因此 可以用map进行优化。
			// 同时 pre是由前一个元素渐次相加，可以用pre待敌preSum[]数组
			count += m[pre-k]
		}
		m[pre]++
	}
	return count
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum3(nums, k))
	fmt.Println(subarraySum(nums, k))

}
