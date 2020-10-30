package main

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	for i := 0; i < n-k; i++ {
		for j := i + 1; j <= i+k; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {
	table := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if idx, ok := table[nums[i]]; ok {
			if i-idx <= k {
				return true
			}
			table[nums[i]] = i
		} else {
			table[nums[i]] = i
		}

	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	table := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if idx, ok := table[nums[i]]; ok && i-idx <= k {
			return true
		}
		table[nums[i]] = i
	}
	return false
}

func main() {

}
