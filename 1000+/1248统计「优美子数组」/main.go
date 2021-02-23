package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	pre, start, end, post := 0, 0, 0, 0
	next := make([]int, 0) // 下一个奇数
	next = append(next, -1)
	res := 0
	for i, v := range nums {
		if v%2 == 1 {
			next = append(next, i)
		}
	}
	next = append(next, len(nums))
	if k+2 > len(next) {
		return 0

	}
	i := k
	for i+1 < len(next) {
		pre = next[i-k]
		start = next[i-k+1]
		end = next[i]
		post = next[i+1]
		res += (start - pre) * (post - end)
		i++

	}
	return res
}

func numberOfSubarrays1(nums []int, k int) int {
	next := make([]int, 0) // 下一个奇数
	next = append(next, -1)
	res := 0
	for i, v := range nums {
		if v&1== 1 {
			next = append(next, i)
		}
	}
	next = append(next, len(nums))
	if k+2 > len(next) {
		return 0

	}
	i := k
	for i+1 < len(next) {
		res += (next[i-k+1] - next[i-k]) * ( next[i+1] - next[i])
		i++

	}
	return res
}
func main() {
	nums := []int{2, 2, 2, 1, 2, 2, 3, 2, 2, 2}
	k := 2
	fmt.Println(numberOfSubarrays1(nums, k))

}
