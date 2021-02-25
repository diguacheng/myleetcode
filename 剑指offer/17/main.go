package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	counts := int(math.Pow(10.0, float64(n)) - 1)
	res := make([]int, counts)
	for i := 0; i < counts; i++ {
		res[i] = i + 1

	}
	return res
}

func printNumbers1(n int) {
	if n <= 0 {
		return
	}
	s := []byte{}
	for i := 0; i < n; i++ {
		s = append(s, '0')
	}
	for !Incerment(s) {
		fmt.Println(string(s))

	}
	return
}

func Incerment(s []byte) bool {
	flag := false  // 标记是否溢出
	nTakeOver := 0 // 进位
	nLenth := len(s)
	for i := nLenth - 1; i >= 0; i-- {
		nSum := int(s[i]-'0') + nTakeOver
		if i == nLenth-1 {
			nSum++
		}
		if nSum >= 10 {
			if i == 0 {
				flag = true // 加法溢出了
			} else {
				nSum -= 10
				nTakeOver = 1
				s[i] = byte('0' + nSum)
			}
		} else {
			s[i] = byte('0' + nSum)
			break

		}

	}
	return flag
}

func main() {
	printNumbers1(4)
}
