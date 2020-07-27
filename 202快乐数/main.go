package main

import "fmt"

func f(num int) int {
	res:=0
	for num>0{
		temp:=num%10
		num=num/10
		res=res+temp*temp
		
	}
	return res

}
func isHappy(n int) bool{
	res:=0
	m:=make(map[int]int)
	res=f(n)
	m[n]=1
	for res!=1&&m[res]<=1{
		res=f(res)
		m[res]++
	}
	if res==1{
		return true
	}
	return false
}

func main() {
	for i:=1;i<1000;i++{
		fmt.Println(i,isHappy(i))
	}
	
}