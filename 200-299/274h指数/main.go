package main

import (
	"fmt"
	"sort")


func hIndex(citations []int) int {
	sort.Slice( citations, func(i, j int)bool{return citations[i] > citations[j]})
	n:=len(citations)
	for i := 0; i < n; i++ {
		if citations[i]<i+1{
			return i
		}
	}
	return n
	

}

func main() {
	c:=[]int{3,0,6,1,5}
	fmt.Println(hIndex(c))

}
