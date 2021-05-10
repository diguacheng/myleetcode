package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	s := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(s) > 0 && nums2[i] > s[len(s)-1] {
			m[s[len(s)-1]] = nums2[i]
			s = s[:len(s)-1]
		}
		s = append(s, nums2[i])
	}
	for len(s) > 0 {
		m[s[len(s)-1]] = -1
		s = s[:len(s)-1]
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = m[nums1[i]]
	}
	return res

}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
