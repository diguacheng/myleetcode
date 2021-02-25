package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	l, r := make([]int, n), make([]int, n)
	l[0] = 1
	r[n-1] = 1
	for i := 1; i < n; i++ {
		l[i] = l[i-1] * nums[i-1]
		r[n-i-1] = r[n-i] * nums[n-i]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = l[i] * r[i]
	}
	return ans

}

func productExceptSelf1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	R := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		R = R * nums[i]
	}
	return ans

}
