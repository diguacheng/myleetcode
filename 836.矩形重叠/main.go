package main

import (
	"fmt"
	
)

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[1]>= rec2[3] ||rec1[2]<=rec2[0]||rec1[3]<=rec2[1]||rec1[0]>=rec2[2])

}
func isRectangleOverlap1(rec1 []int, rec2 []int) bool {
	return min(rec1[2],rec2[2])>max(rec1[0],rec2[0])&&min(rec1[3],rec2[3])>max(rec1[1],rec2[1])
	

}

func max(a, b int) int {
	if a>b{
		return a
	}
	return b
}

func min(a, b int) int {
	if a>b{
		return b
	}
	return a

}


func main(){
	rec1 :=[]int {0,0,2,2} 
	rec2 := []int{1,1,3,3}
	fmt.Println(isRectangleOverlap1(rec1,rec2))
	
}