package main

import (
	"fmt"
	"sort"
)


func sortByBits(arr []int) []int {
	table := map[int]int{}
	for _, num := range arr {
		count:=getOnes(num)
		table[num]=count
	}
	fmt.Println(table)
	sort.Slice(arr,func(i,j int)bool{
		return (table[arr[i]]<table[arr[j]])||((table[arr[i]]==table[arr[j]])&&(arr[i]<arr[j]))

	})
	return arr


}

func getOnes(num int) int {
	count:=0 
	for num!=0{
		if num%2==1{
			count++
		}
		num=num>>1
	}
	return count
}

func main() {
	arr:= []int{8,7,6,5,4,3,2,1,0}
	fmt.Println(sortByBits(arr))
	fmt.Println(1e4)
}