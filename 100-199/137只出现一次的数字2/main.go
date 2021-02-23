package main

import "fmt"

func singleNumber1(nums []int) int {
	set := make(map[int]int)
	sum1, sum2 := 0, 0
	for _, num := range nums {
		sum2 += num
		if _, ok := set[num]; !ok {
			set[num] = 1
			sum1 += num

		}
	}
	return (3*sum1 - sum2) / 2
}

func singleNumber2(nums []int) int {
	hashmap := make(map[int]int)
	for _, num := range nums {
		hashmap[num]++
	}
	for k, count := range hashmap {
		if count == 1 {
			return k

		}
	}
	return 0
}

func singleNumber3(nums []int) int {
	res := [64]int{}
	for _, num := range nums {
		for i := 0; i < 64; i++ {
			res[i] += num & 1
			num = num >> 1
		}
	}
	r := 0
	for i := 0; i < 64; i++ {
		res[i] = res[i] % 3
		r += res[i] << i
	}
	return r
}

func singleNumber(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = ^two & (one ^ num)
		two = ^one & (two ^ num)
	}
	return one

}

func main() {
	nums := []int{2, 2, -3, 2}
	fmt.Println(singleNumber3(nums))

}
