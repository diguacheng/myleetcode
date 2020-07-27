package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	// 枚举法
	l1 := len(str1)
	l2 := len(str2)
	min := l1
	if l1 > l2 {
		min = l2
	}
	for i := min; i >= 1; i-- {
		if l1%i == 0 && l2%i == 0 {
			temp := str1[:i]
			if check(temp, str1) && check(temp, str2) {
				return temp
			}
		}

	}
	return ""

}

func check(t, s string) bool {
	lenx := len(s) / len(t)
	ans := ""
	for i := 1; i <= lenx; i++ {
		ans = ans + t
	}
	return ans == s
}

func gcdOfStrings1(str1 string, str2 string) string {
	// 枚举优化

	l1 := len(str1)
	l2 := len(str2)

	res := str1[:gcd(l1, l2)]
	if check(res, str1) && check(res, str2) {
		return res
	}
	return ""

}

func gcd(a, b int) int {
	temp := 0
	for {
		temp = a % b
		if temp > 0 {
			a = b
			b = temp
		} else {
			return b
		}
	}
}

func gcdOfStrings2(str1 string, str2 string) string {
	// 数学
	if str1+str2!=str2+str1{
		return ""
	}
	l1 := len(str1)
	l2 := len(str2)
	res := str1[:gcd(l1, l2)]
	return res


}

func main() {
	str1 := "LEET"
	str2 := "CODE"
	fmt.Println(gcdOfStrings(str1, str2))

}
