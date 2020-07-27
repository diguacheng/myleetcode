package main

import (
	"fmt"
)

// 正则表达式匹配
func isMatch(s string, p string) bool {
	//fmt.Println(s, p, len(p))
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}
		return false
	}
	if len(p) == 1 {
		if len(s) == 0 {
			return false
		}
		if s[0] == p[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		}
		return false
	}
	if p[1] == '*' {
		if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
			return isMatch(s[1:], p[2:]) || isMatch(s[1:], p[:]) || isMatch(s, p[2:])
		}
		return isMatch(s, p[2:])
	}
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		return isMatch(s[1:], p[1:])
	}
	return false

}

func isMatch3(s string, p string) bool {
	flag := false
	if len(p) == 0 {
		if len(s) == 0 {
			flag = true
		} else {
			flag = false
		}

	}
	if len(p) == 1 {
		if len(s) == 0 {
			flag = false
		} else if s[0] == p[0] || p[0] == '.' {
			flag = isMatch(s[1:], p[1:])
		} else {
			flag = false
		}
	}
	if len(p) > 1 {
		if p[1] == '*' {
			if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
				flag = isMatch(s[1:], p[2:]) || isMatch(s[1:], p[:]) || isMatch(s, p[2:])
			} else {
				flag = isMatch(s, p[2:])
			}

		} else if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
			flag = isMatch(s[1:], p[1:])
		}

	}

	return flag

}

func isMatch4(s string, p string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(p)+1)
	}
	return isMatchCore(s, p, 0, 0, dp)
}
func isMatchCore(s, p string, i, j int, dp [][]int) bool {
	if dp[i][j] == 1 {
		return true
	} else if dp[i][j] == -1 {
		return false
	}
	flag := false
	if j == len(p) {
		if i == len(s) {
			flag = true
		} else {
			flag = false
		}

	}
	if j == len(p)-1 {
		if i == len(s) {
			flag = false
		} else if s[i] == p[j] || p[j] == '.' {
			flag = isMatchCore(s, p, i+1, j+1, dp)
		} else {
			flag = false
		}
	}
	if j < len(p)-1 {
		if p[j+1] == '*' {
			if i <= len(s)-1 && (s[i] == p[j] || p[j] == '.') {
				flag = isMatchCore(s, p, i+1, j+2, dp) || isMatchCore(s, p, i+1, j, dp) || isMatchCore(s, p, i, j+2, dp)
			} else {
				flag = isMatchCore(s, p, i, j+2, dp)
			}

		} else if i <= len(s)-1 && (s[i] == p[j] || p[j] == '.') {
			flag = isMatchCore(s, p, i+1, j+1, dp)
		}

	}
	if flag == true {
		dp[i][j] = 1
	} else {
		dp[i][j] = -1

	}
	return flag

}

func isMatch2(s string, p string) bool {
	//fmt.Println(s, p)
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) != 0 && len(p) == 0 {
		return false
	}
	if len(p) >= 2 {
		if p[1] == '*' {
			if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
				return isMatch2(s[1:], p[2:]) || isMatch2(s[1:], p[:]) || isMatch2(s, p[2:])
			}
			return isMatch2(s, p[2:])
		}
	}
	if len(s) == 0 {
		return false
	}
	if s[0] == p[0] || p[0] == '.' && len(s) != 0 {
		return isMatch2(s[1:], p[1:])
	}

	return false

}

func main() {
	s := "aa"
	p := "a*"

	fmt.Println(isMatch(s, p))
	fmt.Println(isMatch4(s, p))

}
