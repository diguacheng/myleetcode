package main

import "fmt"

func characterReplacement(s string, k int) int {
	window := make(map[byte]int)
	left, right := 0, 0
	maxL := 0
	n := len(s)

	check := func() bool {
		
		maxValue := 0
		for _, v := range window {
		
			if v > maxValue {
				maxValue = v
			}
		}
		return right-left-maxValue <= k
	}
	for right < n {
		window[s[right]]++
		right++
		if check() {
			if right-left > maxL {
				maxL = right - left
			}
		} else {
			window[s[left]]--
			left++
		}

	}
	return maxL
}

func characterReplacement2(s string, k int) int {
	window := [26]int{}
	left, right := 0, 0
	maxL := 0
	n := len(s)

	check := func() bool {
		
		maxValue := 0
		for _, v := range window {
		
			if v > maxValue {
				maxValue = v
			}
		}
		return right-left-maxValue <= k
	}
	for right < n {
		window[s[right]-'A']++
		right++
		if check() {
			if right-left > maxL {
				maxL = right - left
			}
		} else {
			window[s[left]-'A']--
			left++
		}

	}
	return maxL
}

func characterReplacement1(s string, k int) int {
	window := make(map[byte]int)
	left, right := 0, 0
	maxL := 0
	n := len(s)

	check := func() bool {
		sum := 0
		maxValue := 0
		for _, v := range window {
			sum += v
			if v > maxValue {
				maxValue = v
			}
		}
		return sum-maxValue > k
	}

	for right = 0; right < n; right++ {
		window[s[right]]++
		if right+1 >= n || s[right+1] != s[right] {
			for check() {
				window[s[left]]--
				left++
			}
			if right-left+1 > maxL {
				maxL = right - left + 1
			}
		}

	}
	return maxL
}


func characterReplacement3(s string, k int) int {
	window := [26]int{}
	left, right := 0, 0
	maxL := 0
	n := len(s)

	check := func() bool {

		maxValue := 0
		for _, v := range window {
	
			if v > maxValue {
				maxValue = v
			}
		}
		return right-left+1-maxValue > k
	}

	for right = 0; right < n; right++ {
		window[s[right]-'A']++
		if right+1 >= n || s[right+1] != s[right] {
			for check() {
				window[s[left]-'A']--
				left++
			}
			if right-left+1 > maxL {
				maxL = right - left + 1
			}
		}

	}
	return maxL
}

// 上面的效率太低，把哈希表map换为数组后，提升了点，但是效率也不高 
// 下面这个双百 
func characterReplacement100(s string, k int) int {
	if s == "" || len(s) == 0 {
		return 0
	}

	hash := make([]int, 26)
	l, maxCount, result := 0, 0, 0
	for r := 0; r < len(s); r++ {
		hash[s[r]-'A']++

		if maxCount < hash[s[r]-'A'] {
			maxCount = hash[s[r]-'A']
		}

		for r-l+1-maxCount > k {
			hash[s[l]-'A']--
			l++
		}

		if result < r-l+1 {
			result = r - l + 1
		}
	}
	return result
}

func main() {
	s := "BAAAB"
	fmt.Println(characterReplacement2(s, 2))

}
