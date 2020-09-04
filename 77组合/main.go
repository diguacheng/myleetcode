package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	temp := make([]int, 0)
	backTrack(1, n, k, temp, &res)
	return res
}

func backTrack(start, end, k int, temp []int, res *[][]int) {
	for i := start; i <= end; i++ {
		temp = append(temp, i)
		if k == 1 {

			ans := make([]int, len(temp))
			copy(ans, temp)
			*res = append(*res, ans)
		} else {
			backTrack(i+1, end, k-1, temp, res)

		}
		temp = temp[:len(temp)-1]
	}
}

func combine1(n int, k int) [][]int {
	nums := make([]int, k+1)
	for i := 1; i < k+1; i++ {
		nums[i-1] = i
	}
	nums[len(nums)-1] = n + 1
	res := make([][]int, 0)
	j := 0
	for j < k {
		ans := make([]int, k)
		copy(ans, nums[:k])
		res = append(res, ans)
		j = 0
		for j < k && nums[j+1] == nums[j]+1 {
			nums[j] = j + 1
			j++
		}
		nums[j]++
	}
	return res
}

func main() {
	fmt.Println(combine1(5, 3))

}
