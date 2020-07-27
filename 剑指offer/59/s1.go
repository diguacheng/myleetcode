package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := make([]int, 0)
	for i := 0; i+k-1 < len(nums); i++ {
		max := nums[i]
		for j := i + 1; j < i+k; j++ {
			if max < nums[j] {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := make([]int, 0)
	if len(nums) >= k && k >= 1 {
		list := make([]int, 0)
		for i := 0; i < k; i++ {
			for len(list) != 0 && nums[i] >= nums[list[len(list)-1]] {
				list = list[:len(list)-1]
			}
			list = append(list, i)
		}
		for i := k; i < len(nums); i++ {
			res = append(res, nums[list[0]])
			for len(list) != 0 && nums[i] >= nums[list[len(list)-1]] {
				list = list[:len(list)-1]
			}
			if len(list) != 0 && list[0] <= i-k {
				list = list[1:]
			}
			list = append(list, i)
		}
		res = append(res, nums[list[0]])

	}
	return res

}

type MaxQueue struct {
	max      []int
	listData []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int, 0),
		make([]int, 0),
	}

}

func (this *MaxQueue) Max_value() int {
	if len(this.listData) == 0 {
		return -1
	}
	return this.max[0]

}

func (this *MaxQueue) Push_back(value int) {
	for len(this.max) != 0 && this.max[len(this.max)-1] < value {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
	this.listData = append(this.listData, value)

}

func (this *MaxQueue) Pop_front() int {
	if len(this.listData) == 0 {
		return -1
	}
	x := this.listData[0]
	if x == this.max[0] {
		this.max = this.max[1:]
	}
	this.listData = this.listData[1:]
	return x

}

func main() {
	nums := []int{1}
	k := 1
	fmt.Println(maxSlidingWindow1(nums, k))

}
