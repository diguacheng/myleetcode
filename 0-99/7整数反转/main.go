package main

import (
	"fmt"
)

func reverse(x int) int {
	res:=0
	for x!=0{
		temp:=x%10
		x=x/10
		if -(1<<31)>res||res>(1<<31-1){
			return 0
		}
		res=res*10 +temp
	}
	return res

}

func main(){
	fmt.Println(reverse(123))
}