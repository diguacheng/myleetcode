package main

import (
	"fmt"
	"sort"
)


type ssort [][]int

func (s ssort) Len() int { return len(s) }
func (s ssort) Less(i, j int) bool {
	if s[i][0] > s[j][0] {
		return true
	}
	if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}
	return false
}



func (s ssort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(ssort(people))
	fmt.Println(people)
	n:=len(people)
	res:=make([][]int,n)
	for i := 0; i < n; i++ {
		res[i]=make([]int,2)
	}
	insert:=func(p []int,i int){
		idx:=i
		for idx!=p[1]{
			res[idx]=res[idx-1]
			idx--
		}
		res[idx]=p
	}
	for i:=0;i<n;i++{
		insert(people[i],i)
	}
	return res
}

func reconstructQueue1(people [][]int) [][]int {
	sort.Slice(people, func(i, j int)bool{
		if people[i][0]>people[j][0]||(people[i][0]==people[j][0]&&people[i][1]<people[j][1]){
			return true
		}
		return false
	})
	fmt.Println(people)

	n:=len(people)

	var p []int
	for i:=0;i<n;i++{
		p=people[i]
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]]=p
		//fmt.Println(people)
		
	}
	return people
}

// type ssort [][]int

// func (s ssort) Len() int { return len(s) }
// func (s ssort) Less(i, j int) bool {
// 	if s[i][1] < s[j][1] {
// 		return true
// 	}
// 	if s[i][1] == s[j][1] {
// 		return s[i][0] < s[j][0]
// 	}
// 	return false
// }



// func (s ssort) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func reconstructQueue(people [][]int) [][]int {
// 	sort.Sort(ssort(people))
// 	n:=len(people)
// 	res:=make([][]int,n)
// 	for i := 0; i < n; i++ {
// 		res[i]=make([]int,2)
// 	}
// 	var p []int
// 	var count int 
// 	var idx int 
// 	for i := 0; i < n; i++ {
// 		p=people[i]
// 		count=0
// 		idx=i
// 		for j:=0;j<i;j++{
// 			if res[j][0]>=p[0]{
// 				count++
// 			}
// 			if count==p[1]+1{
// 				for idx!=j{
// 					res[idx]=res[idx-1]
// 					idx--
// 				}
// 				break
// 			}
// 		}
// 		res[idx]=p
// 	}
// 	return res
// }
func main() {
	people:=[][]int{
		{9,0},{7,0},{1,9},{3,0},{2,7},{5,3},{6,0},{3,4},{6,2},{5,2},
	}
	fmt.Println( reconstructQueue(people))

}