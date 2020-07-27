package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	
	str:=strconv.Itoa(num)
	n:=len(str)
	counts:=make([]int,n)
	count:=0
	for i:=n-1; i>=0; i-- {
		count=0
		if i<n-1{
			count=counts[i+1]
		}else{
			count=1
		}
		if i<n-1{
			number,_:=strconv.Atoi(str[i:i+2])
			if number>=10&&number<=25{
				if i<n-2{
					count+=counts[i+2]
				}else{
					count++
				}
			}

		}
		counts[i]=count
	}
	count=counts[0]
	return count


}

func main() {
	num:=122
	fmt.Println(translateNum(num))

}