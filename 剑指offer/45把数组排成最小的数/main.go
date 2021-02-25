package main

import (
	"fmt"
	"sort"
	"strconv"
)

type intstring []int

func (s intstring) Len() int { return len(s) }
func (s intstring) Less(i, j int) bool {
	a := strconv.Itoa(s[i])
	b := strconv.Itoa(s[j])
	var n int
	if len(a) < len(b) {
		n = len(b)
	} else {
		n = len(a)
	}
	for i := 0; i <= n; i++ {
		idxa := i % len(a)
		idxb := i % len(b)
		if a[idxa] < b[idxb] {
			return true
		} else if a[idxa] > b[idxb] {
			return false
		}

	}
	return false
}

func (s intstring) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func minNumber(nums []int) string {

	sort.Sort(intstring(nums))
	//fmt.Println(nums)
	s := ""
	for _, num := range nums {
		s = s + strconv.Itoa(num)

	}
	return s

}

func minNumber1(nums []int) string {
	numstr := []string{}
	for _, num := range nums {
		numstr = append(numstr, strconv.Itoa(num))
	}
	sort.Slice(numstr, func(i, j int) bool {
		a := numstr[i]
		b := numstr[j]
		var n int
		if len(a) < len(b) {
			n = len(b)
		} else {
			n = len(a)
		}
		for i := 0; i <= n; i++ {
			idxa := i % len(a)
			idxb := i % len(b)
			if a[idxa] < b[idxb] {
				return true
			} else if a[idxa] > b[idxb] {
				return false
			}

		}
		return false

	})

	s := ""
	for _, num := range numstr {
		s = s + num

	}
	return s

}

func minNumber2(nums []int) string {
	numstr := []string{}
	for _, num := range nums {
		numstr = append(numstr, strconv.Itoa(num))
	}
	sort.Slice(numstr, func(i, j int) bool {
		a, _ := strconv.Atoi(numstr[i] + numstr[j])
		b, _ := strconv.Atoi(numstr[j] + numstr[i])
		return a < b

	})

	s := ""
	for _, num := range numstr {
		s = s + num

	}
	return s

}

func main() {
	var nums intstring
	nums = []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}
	//nums=[]int{2,10}
	fmt.Println(minNumber1(nums))
	fmt.Println(minNumber(nums))

}
