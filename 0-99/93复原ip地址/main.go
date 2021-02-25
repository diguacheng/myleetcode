package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 需要提前做出一些判断。
func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}
	res := make([]string, 0)
	temp := make([]byte, 0)
	backTrack(s, temp, 4, &res)
	return res

}

func backTrack(s string, temp []byte, k int, res *[]string) {
	n := len(s)
	if n < k || n > 3*k {
		return
	}
	if k == 0 && n == 0 {
		ans := make([]byte, len(temp)-1)
		copy(ans, temp[:len(temp)-1])
		*res = append(*res, string(ans))
		return
	}

	if s[0] == '0' {
		temp = append(temp, s[0])
		temp = append(temp, '.')
		backTrack(s[1:], temp, k-1, res)
		temp = temp[:len(temp)-2]
		return
	}
	for i := 0; i < n && i <= 2; i++ {
		num, _ := strconv.Atoi(s[:i+1])
		if num >= 0 && num <= 255 {
			temp = append(temp, s[:i+1]...)
			temp = append(temp, '.')
			backTrack(s[i+1:], temp, k-1, res)
			temp = temp[:len(temp)-i-2]
		}
	}

}

func restoreIpAddresses1(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}
	res := make([]string, 0)
	temp := make([]string, 0)
	backTrack1(s, temp, 4, &res)
	return res

}

func backTrack1(s string, temp []string, k int, res *[]string) {
	n := len(s)
	if n < k || n > 3*k {
		return
	}
	if k == 0 && n == 0 {
		ans := make([]string, len(temp))
		copy(ans, temp)
		str := strings.Join(ans, ".")
		*res = append(*res, str)
		return
	}
	if s[0] == '0' {
		temp = append(temp, s[0:1])
		backTrack1(s[1:], temp, k-1, res)
		temp = temp[:len(temp)-1]
		return
	}
	for i := 0; i < n && i <= 2; i++ {
		num, _ := strconv.Atoi(s[:i+1])
		if num >= 0 && num <= 255 {
			temp = append(temp, s[:i+1])
			backTrack1(s[i+1:], temp, k-1, res)
			temp = temp[:len(temp)-1]
		}
	}

}

func main() {
	s := "101023"
	fmt.Println(restoreIpAddresses(s))
	fmt.Println(restoreIpAddresses1(s))

}
