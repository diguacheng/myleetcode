package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	len2 := len(s2)
	len1 := len(s1)
	if len2 < len1 {
		return false
	}
	need := [26]int{}
	window := [26]int{}
	for i := 0; i < len1; i++ {
		need[s1[i]-'a']++
	}
	check := func() bool {
		for i := 0; i < 26; i++ {
			if window[i] != need[i] {
				return false
			}
		}
		return true
	}
	match := 0 // 表示匹配的字符个数
	left, right := 0, 0
	var idx, leftidx byte
	for right < len2 {
		idx = s2[right] - 'a'
		window[idx]++
		right++
		if window[idx] <= need[idx] {
			match++
		}
		for match == len1 {
			if check() {
				return true
			}
			leftidx = s2[left] - 'a'
			left++
			if window[leftidx] <= need[leftidx] {
				match--
			}
			window[leftidx]--
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidboaoo"
	fmt.Println(checkInclusion(s1,s2))
}