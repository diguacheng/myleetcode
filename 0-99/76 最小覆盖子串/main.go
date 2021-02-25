package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]

}

func minWindow1(s string, t string) string {
	need := make(map[byte]int)   //记录字符串t的各个元素
	window := make(map[byte]int) // 记录窗口内字符的个数
	lent := len(t)
	for i := 0; i < lent; i++ {
		need[t[i]]++
	}
	lens := len(s)
	distance := 0
	left, right := 0, 0 // 窗口的左右边界 左闭右开
	begin, end := 0, 0
	minlen := lens + 1 // 最小字串的长度
	for right < lens {
		if need[s[right]] == 0 {
			right++
			continue
		}
		if window[s[right]] == need[s[right]] {
			distance++
		}
		window[s[right]]++
		right++
		for distance == len(need) {
			if right-left < minlen {
				minlen = right - left
				begin = left
				end = right
			}
			if need[s[left]] == 0 {
				left++
				continue
			}
			if window[s[left]] == need[s[left]] {
				distance--
			}
			window[s[left]]--
			left++
		}

	}
	if minlen == lens+1 {
		return ""
	}
	return s[begin:end]

}

func minWindow2(s string, t string) string {
	need := make(map[byte]int)   //记录字符串t的各个元素
	window := make(map[byte]int) // 记录窗口内字符的个数
	lent := len(t)
	for i := 0; i < lent; i++ {
		need[t[i]]++
	}
	lens := len(s)
	distance := 0
	left, right := 0, 0 // 窗口的左右边界 左闭右开
	begin, end := 0, 0
	minlen := lens + 1 // 最小字串的长度
	for right < lens {
		if need[s[right]] == 0 {
			right++
			continue
		}
		rightchar := s[right]
		if window[rightchar] < need[rightchar] {
			distance++
		}
		window[rightchar]++
		right++

		for distance == lent {
			if right-left < minlen {
				minlen = right - left
				begin = left
				end = right
			}
			if need[s[left]] == 0 {
				left++
				continue
			}
			leftchar := s[left]
			if window[leftchar] == need[leftchar] {
				distance--
			}
			window[leftchar]--
			left++
		}

	}
	if minlen == lens+1 {
		return ""
	}
	return s[begin:end]

}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow1(s, t))

}
