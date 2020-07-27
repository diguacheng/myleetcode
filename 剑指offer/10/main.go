package main

import (
	"fmt"
)

func fib(n int) int {
	temp:=[]int{0,1}[:]
	/*
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	经过分析，发现 上述代码 可以由最后的return 包含住 
	*/
	for i:=2;i<=n;i++{
		temp=append(temp,(temp[i-1]+temp[i-2])%1000000007)
	}
	return temp[n]

}

func main(){
	fmt.Println(fib(95))

	

}