package main

func maxProduct(nums []int) int {

	minF, maxF, res := nums[0], nums[0], nums[0]
	var tempmin, tempmax int
	for i := 1; i < len(nums); i++ {
		tempmin, tempmax = minF, maxF
		maxF = max(max(tempmin*nums[i], tempmax*nums[i]), nums[i])
		minF = min(min(tempmin*nums[i], tempmax*nums[i]), nums[i])
		res = max(maxF, res)

	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

}
