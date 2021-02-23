package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	table := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	var f func(temp []byte, digits string)
	f = func(temp []byte, digits string) {
		if len(digits) == 0 {
			res = append(res, string(temp))
			return
		}
		char := digits[0]
		str := table[int(char-'2')]
		for _, v := range str {
			temp = append(temp, byte(v))
			f(temp, digits[1:])
			temp = temp[:len(temp)-1]
		}
	}
	f([]byte{}, digits)
	return res

}


func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	table := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	var f func(temp *[]byte, digits string)
	f = func(temp *[]byte, digits string) {
		if len(digits) == 0 {
			s := make([]byte, len(*temp))
			copy(s, *temp)
			res = append(res, string(s))
			return

			// 因为用的是指针变量，这里应该new一个新的变量 并将变量复制归来，或者不要用指针。
		}
		char := digits[0]
		str := table[int(char-'2')]
		for _, v := range str {
			*temp = append(*temp, byte(v))
			f(temp, digits[1:])
			*temp = (*temp)[:len(*temp)-1]
		}
	}
	f(&[]byte{}, digits)
	return res

}
func main() {
	fmt.Println(letterCombinations1("234"))

}
