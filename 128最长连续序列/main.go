package main

import "fmt"

func longestConsecutive(nums []int) int {
	// 哈希法
	m:=make(map[int]bool)
	for _,num:=range nums{
		m[num] = true
	}
	ans:=0
	for num:=range m{
		if !m[num-1] {
			curnum:=num
			cunrans:=0
			for m[curnum]{
				curnum++
				cunrans++
			}
			if cunrans>ans{
				ans=cunrans
			}
		}
	}
	return ans

}

func longestConsecutive1(nums []int) int {
	// 并查集思想
	ans:=0
	m:=make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 分组
		m[nums[i]]=i+1
	}
	for num:=range m{
		// 合并组
		temp:=num
		for	m[temp-1]>0{
			temp=temp-1
		}
		m[num]=m[temp]
	}
	l:=make(map[int]int)
	for _,v:=range m{
		// 找出个数最大的组
		l[v]++
		if l[v]>ans{
			ans=l[v]
		}
	}
	return ans

}

func main(){
	nums:=[]int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive1(nums))

}