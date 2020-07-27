package main

import "fmt"

func firstUniqChar(s string) byte {
	if len(s)==0{
		return ' '
	}
	n:=len(s)
	m:=make(map[byte]int)
	for i := 0; i < n; i++ {
		m[s[i]]++
	}
	for i := 0; i < n; i++ {
		if m[s[i]]==1{
			return s[i]
		}
	}
	return ' '

}

func firstUniqChar1(s string) byte {
	if len(s)==0{
		return ' '
	}
	list:=[256]int{}
	n:=len(s)
	
	for i := 0; i < n; i++ {
		list[s[i]]++
	}
	for i := 0; i < n; i++ {
		if list[s[i]]==1{
			return s[i]
		}
	}
	return ' '

}




func main() {
	fmt.Println(firstUniqChar("abaccdeff"))

}