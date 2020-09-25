package main

import (
	"fmt"	
	"sort"
)


func leastInterval(tasks []byte, n int) int {
	l:=len(tasks)
	list:=[26]int{}
	for i:=0;i<l;i++ {
		list[tasks[i]-'A']++
	}
	sort.Ints(list[:]) // 升序排列 
	time:=0
	i:=0
	for list[25]>0{
		i=0
		for i<=n{ 
			 // 以n+1 为一轮，在这一轮，前n+1大的元素减一 
			if list[25]==0{
				break
			}
			if i<26&&list[25-i]>0{
				list[25-i]--
			}
			time++
			i++
		}
		sort.Ints(list[:])
	}
	return time

}



func leastInterval1(tasks []byte, n int) int {
	l:=len(tasks)
	list:=[26]int{}
	for i:=0;i<l;i++ {
		list[tasks[i]-'A']++
	}
	sort.Ints(list[:]) // 升序排列 
	maxValue:=list[25]-1
	a:=maxValue*n
	for i:=24;i>=0;i--{
		a-=min(list[i],maxValue)
	}
	if a>0{
		return a+l
	}
	return l

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b 
}

func main() {
	s := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	fmt.Println(leastInterval(s,2))
	fmt.Println(leastInterval1(s,2))
}