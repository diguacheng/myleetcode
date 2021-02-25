package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	list1 := strings.Split(version1, ".")
	list2 := strings.Split(version2, ".")
	var i int
	for i = 0; i < len(list1) && i < len(list2); i++ {
		num1, _ := strconv.Atoi(list1[i])
		num2, _ := strconv.Atoi(list2[i])
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	for i < len(list1) {
		num1, _ := strconv.Atoi(list1[i])
		if num1 > 0 {
			return 1
		}
		i++
	}
	for i < len(list2) {
		num2, _ := strconv.Atoi(list2[i])
		if num2 > 0 {
			return -1
		}
		i++
	}
	return 0

}

func main() {

}
