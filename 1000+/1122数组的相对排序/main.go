package main

import (
	"fmt"
	"sort")


func relativeSortArray(arr1 []int, arr2 []int) []int {
	n := len(arr2)
	table := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		table[arr2[i]] = i
	}
	sort.Slice(arr1,func(i,j int)bool{
		x,y:=arr1[i],arr1[j]
		px,hasx:=table[x]
		py,hasy:=table[y]
		if hasx&&hasy{
			return px<py
		}
		if !hasx&&!hasy{
			return x<y
		}
		if !hasx{
			return false
		}
		return true

	})
	return arr1
}

func main() {
	arr1:=[]int{2,3,1,3,2,4,6,0,9,2,19}
	arr2:=[]int{2,1,4,3,9,6}
	fmt.Println(relativeSortArray(arr1,arr2))

}