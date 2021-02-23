package main

import (
	"fmt"
	"strconv"
)

func lengthOfLongestSubstring(s string) int {
	end,res:=-1,0// 右指针
	m:=make(map[byte]int)
	n:=len(s)
	for i:=0;i<n;i++{
		if i!=0{
			//左指针向右移动，移除一个字符
			delete(m,s[i-1])
		}
		for end+1<n&&m[s[end+1]]==0{
			m[s[end+1]]++
			end++
		}
		if res<end-i+1{
			res=end-i+1		
		}

	}
	return res
}

func lengthOfLongestSubstring1(s string) int {
	list:=make([]int,128)
	res:=0
	start,end:=0,-1
	n:=len(s)
	for start<n{
		if end+1<n&&list[s[end+1]]==0{
			end++
			list[s[end]]++
		}else{
			list[s[start]]--
			start++
		}
		if res<end-start+1{
			res=end-start+1
		}
	}
	return res
}

func main(){
	
	fmt.Println( strconv.Itoa(1<<63-1) )

}