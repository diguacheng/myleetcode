package main

import (
	"fmt"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	flag := 1
	index := -1
	// 找到数字的开头
	for i := 0; i < len(str); i++ {
		if str[i] == '+' {
			flag = 1
			index = i + 1
			break

		}
		if str[i] == '-' {
			flag = 0
			index = i + 1
			break
		}
		if str[i] >= '0' && str[i] <= '9' {
			index = i
			break

		}
		if str[i] != ' ' {
			return 0
		}
	}
	// 没有符合条件的数字开头 返回0
	if index == -1 {
		return 0
	}
	sum := 0
	v := 1 << 31 // 2**31
	min := -v
	max := v - 1
	// 转换成数字 注意防止溢出
	for i := index; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			sum = sum*10 + int(str[i]-'0')
			// 这里有一个技巧， 如flag=0 sum>1<<31 如flag=10 sum>1<<31-1, 满足正负数溢出的不同条件
			if sum > v-flag {
				if flag == 0 {
					return min
				}
				return max
			}

		} else {
			break
		}
	}
	// 如果flag=0说明改数是符号
	if flag == 0 {
		sum = -sum
	}
	return sum

}

func main() {
	fmt.Println(myAtoi("9223372036854775808"))

}
