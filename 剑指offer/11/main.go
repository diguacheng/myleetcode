package main

import (
	"fmt"
)

func minArray(numbers []int) int {
	start,end := 0, len(numbers) - 1
	for start<end{
		mid:=(start+end)/2
		if numbers[mid]>numbers[end]{
			start=mid+1
		}else if numbers[mid]<numbers[start]{
			end=mid
		}else{
			end--
		}
	}
	return numbers[start]
}

func main() {
	numbers:=[]int{1,3,5}
	fmt.Println(minArray(numbers))

}