package main

import "fmt"

func strStr0(haystack string, needle string) int {
	// 基本暴力法
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	res := -1
	f := false
	n := len(haystack)
	for i := 0; i < n; i++ {
		if haystack[i] == needle[0] {
			j, k := i+1, 1
			for ; j < n && k < len(needle); j, k = j+1, k+1 {
				if haystack[j] != needle[k] {
					break
				}
			}
			if k >= len(needle) {
				f = true

			}

		}
		if f == true {
			res = i
			return res
		}
	}
	return res

}

func strStr(haystack string, needle string) int {
	// 双指针
	lenh := len(haystack)

	lenn := len(needle)
	if lenh < lenn {
		return -1
	}
	if lenn == 0 {
		return 0
	}
	start, end := 0, 0
	k := 0
	for start < lenh {
		if haystack[start] == needle[k] {
			end = start + 1
			k++
			for end < lenh && k < lenn && haystack[end] == needle[k] {
				k++
				end++
			}
			if k >= lenn {
				return start
			}
			start = end
		}
	}
	return 1

}

func main() {
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr0(haystack, needle))

}
