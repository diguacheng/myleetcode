package main

import "fmt"

func constructArr(a []int) []int {
	n:=len(a)
	if n==0{
		return nil
	}
	x1:=make([]int,n)
	x2:=make([]int,n)
	x1[0]=1
	for i := 1; i < n; i++ {
		x1[i]=x1[i-1]*a[i-1]
	}
	x2[n-1]=1
	for i:=n-2;i>=0;i--{
		x2[i]=x2[i+1]*a[i+1]
	}
	res:=make([]int,n)
	for i := 0; i < n; i++ {
		res[i]=x1[i]*x2[i]
	}
	return res

}

func main() {
	nums:=[]int{1,2,3,4,5}
	fmt.Println(constructArr(nums))

}