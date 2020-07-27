package main

import (
	"fmt"

)

func reverseLeftWords(s string, n int) string {
	// 用切片
	return s[n:]+s[:n]
}

func reverseLeftWords1(s string, n int) string {
	// 若不允许使用切边
	res:=""
	for i:=n;i<len(s);i++{
		res+=string(s[i])
	}
	for i:=0;i<n;i++{
		res+=string(s[i])

	}
	return res
}

func main() {
	s := "lrloseumgh"
	fmt.Println(reverseLeftWords1(s,6))

}
