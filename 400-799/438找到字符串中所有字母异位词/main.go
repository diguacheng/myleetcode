package main

import "fmt"

func findAnagrams(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)
	if lenS < lenP {
		return []int{}
	}
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < lenP; i++ {
		need[p[i]]++
	}
	check := func() bool {
		for char, num := range window {
			if need[char] != num {
				return false
			}
		}
		return true
	}
	match := 0 // 表示匹配的字符种类数
	left, right := 0, 0
	res := []int{}
	var char, leftChar byte
	for right < lenS {
		char = s[right]
		window[char]++
		right++
		if window[char] == need[char] {
			match++
		}
		for match == len(need) {
			if check() {
				res = append(res, left)
			}
			leftChar = s[left]
			left++
			// 这里要注意一下细节,如果need[leftChar]==0，也可能导致match被减
			if need[leftChar] != 0 && need[leftChar] == window[leftChar] {
				match--
			}
			window[leftChar]--
		}
	}
	return res
}

func findAnagrams1(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)
	if lenS < lenP {
		return []int{}
	}
	need := [26]int{}
	window := [26]int{}
	for i := 0; i < lenP; i++ {
		need[p[i]-'a']++
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
	res := []int{}
	var idx, leftidx byte
	for right < lenS {
		idx = s[right] - 'a'
		window[idx]++
		right++
		if window[idx] <= need[idx] {
			match++
		}
		for match == lenP {
			if check() {
				res = append(res, left)
			}
			leftidx = s[left] - 'a'
			left++
			// 这里要注意一下细节,如果need[leftChar]==0，也可能导致match被减
			if window[leftidx] <= need[leftidx] {
				match--
			}
			window[leftidx]--
		}
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams1(s, p))
}
