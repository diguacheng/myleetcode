package main

func findPeakElement(nums []int) int {
	s, e := 0, len(nums)-1
	for s < e {
		mid := (s + e) / 2
		if nums[mid] > nums[mid+1] {
			e = mid
		} else {
			s = mid
		}
	}
	return s

}

func main() {

}
