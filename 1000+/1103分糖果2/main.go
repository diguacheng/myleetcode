package main

import "fmt"
func distributeCandies(candies int, num_people int) []int {//暴力法

	res:=make([]int, num_people)
	i:=0
	for candies>0{
		if candies<i+1{
			res[i%num_people]=res[i%num_people]+candies
			candies=0
			break

		}
		res[i%num_people]=res[i%num_people]+i+1
		candies=candies-i-1
		i++
	}
	return res


}
func distributeCandies1(candies int, num_people int) []int {
	//数学法 
	// 目前还不会 

	res:=make([]int, num_people)
	i:=0
	for candies>0{
		if candies<i+1{
			res[i%num_people]=res[i%num_people]+candies
			candies=0
			break

		}
		res[i%num_people]=res[i%num_people]+i+1
		candies=candies-i-1
		i++
	}
	return res


}



func main() {
	candies := 7
	num_people := 4
	fmt.Println(distributeCandies(candies,num_people))

}