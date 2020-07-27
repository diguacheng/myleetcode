package main

func lengthOfLongestSubstring(s string) int {
	list := make([]int, 128)
	n := len(s)
	start, end := 0, -1
	res := 0
	for start < n {
		if end+1 < n && list[s[end+1]] == 0 {
			end++
			list[s[end]]++
		} else {
			list[s[start]]--
			start++
		}
		if res < end-start+1 {
			res = end - start + 1
		}

	}
	return res
}

func main() {

}
