package main

import "fmt"
 // 关键是找到这两个二进制的公共前缀 
//2147483647  = 1<<31-1 
func rangeBitwiseAnd(m int, n int) int {
	count:=0
	for m!=n{
		m=m>>1
		n=n>>1
		count++
	}
	return m<<count
}


func rangeBitwiseAnd1(m int, n int) int {
	for m<n{
		n=n&(n-1)
	}
	return m&n
}

func main(){
	fmt.Println(rangeBitwiseAnd(4,7))
	fmt.Println(rangeBitwiseAnd(1,2))
}