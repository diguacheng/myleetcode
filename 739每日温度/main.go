package main

import "fmt"

func dailyTemperatures(T []int) []int {
	// 暴力法 
	res:=make([]int,len(T))
	start,end := 0,0
	for start<len(T){
		end=start+1
		for end<len(T){
			if T[start]<T[end] {
				res[start]=end-start
				break
			}
			end++
		}
		start++
	}
	return res

}

func dailyTemperatures1(T []int) []int {
	res:=make([]int,len(T))
	stack:=make([]int,0)
	var pre int
	for i:=0;i<len(T);i++ {
		if len(stack)==0{
			stack=append(stack,i)
			continue
		}
		pre=stack[len(stack)-1]
		if T[i]<=T[pre]{
			stack=append(stack,i)
			continue
		}
		for T[i]>T[pre] {
			res[pre]=i-pre
			stack=stack[:len(stack)-1]
			if len(stack)==0{
				break
			}
			pre=stack[len(stack)-1]
		}
		stack=append(stack,i)
	}
	return res

}

func dailyTemperatures2(T []int) []int {
	res:=make([]int,len(T))
	stack:=make([]int,0)
	var pre int
	for i:=0;i<len(T);i++ {
		for len(stack)>0&&T[i]>T[stack[len(stack)-1]]{
			pre=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			res[pre]=i-pre
			
		}
		stack=append(stack,i)
	}
	return res

}
func main(){
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
	fmt.Println(dailyTemperatures2(temperatures))


}