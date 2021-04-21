package main

import "fmt"

// 暴力不可行 
func isScramble1(s1 string, s2 string) bool {
	if len(s1) == 1 {
		return s1==s2
	}
	ans := DFS(s1)
	return  ans[s2] == 1
}

func DFS(s string) map[string]int {
	if len(s) == 1 {
		return map[string]int{s: 1}
	}
	n := len(s)
	res := make(map[string]int)
	for i := 1; i < n; i++ {
		arr1 := DFS(s[:i])
		arr2 := DFS(s[i:])
		for a := range arr1 {
			for b := range arr2 {
				res[a+b] = 1
				res[b+a] = 1
			}
		}
	}
	return res
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) == 1 {
		return s1==s2
	}
	n:=len(s1)
	m1:=make(map[byte]int)
	m2:=make(map[byte]int)
	for i:=0;i<n;i++{
		m1[s1[i]]++
		m2[s2[i]]++
		if check(m1,m2){
			if isScramble(s1[i+1:],s2[i+1:]){
				return true
			}
		}
	}
	return false

}

func check(m1,m2 map[byte]int)bool{
	for k,v:=range m1{
		if m2[k]!=v{
			return false
		}
	}
	return true
}



func main() {
	s1 := "abcdefghijklmn"
	s2 := "efghijklmncadb"
	fmt.Println(isScramble1(s1,s2))
	fmt.Println(isScramble(s1,s2))
}