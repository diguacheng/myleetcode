package main

import "fmt"

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	d := []int{2, 3, 5}

	for num != 1 {
		flag := false
		for i := 0; i < 3; i++ {
			if num%d[i] == 0 {
				num = num / d[i]
				flag = true
				break
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUgly(-2147483648))
}
