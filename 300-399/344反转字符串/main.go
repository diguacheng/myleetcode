package main

import "fmt"

func reverseString(s []byte) {
	n := len(s)
	l := n / 2
	var temp byte
	for i := 0; i < l; i++ {
		temp = s[i]
		s[i] = s[n-1-i]
		s[n-1-i] = temp
	}
}

func reverseString1(s []byte) {
	n := len(s)
	start, end := 0, n-1
	var temp byte
	for start < end {
		temp = s[start]
		s[start] = s[end]
		s[end] = temp
		start++
		end--
	}
}

func main() {
	s := []byte{'a', 'b', 'c', 'd'}
	reverseString(s)
	fmt.Println(string(s))

}
