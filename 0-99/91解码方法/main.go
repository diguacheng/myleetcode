package main

import (
	"fmt"
)


// dp[i表示0..i表示的数 
// 若 s[i]==0 且s[i-1]=='1'||s[i-1]=='2' dp[i]==dp[i-2]   否则为0    比如 10
// 若s[i]!=0,且 s[i-1]==1||(s[i-1]=='2'&&s[i]<='6') dp[i]=dp[i-1]+dp[i-2] 比如 12 23 等 反之 do[i]=p[i-1] 比如 36 29 等 

func numDecodings(s string) int {
	n := len(s)
	if n==0||s[0]=='0' {
		return 0
	}
	if n==1{
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
 
	for i := 1 ;i<n;i++{
		if s[i]=='0'{
			if s[i-1]!='1'&&s[i-1]!='2'{
				return 0
			}
			if i>1{
				dp[i]=dp[i-2]
			}else{
				dp[i]=1
			}
			
		}else if s[i-1]=='1'||(s[i-1]=='2'&&s[i]<='6'){
			if i>1{
				dp[i]=dp[i-2]+dp[i-1]
			}else{
				dp[i]=dp[i-1]+1
			}
		}else{
			dp[i]=dp[i-1]
		}

	}
	return dp[n-1]
}


func numDecodings1(s string) int {
	n := len(s)
	if n==0||s[0]=='0' {
		return 0
	}
	if n==1{
		return 1
	}
	cur,pre:=1,1 // pre 代表dp[i-2]cur  在没更新前 代表 dp[i-1]
	var temp int 
 
	for i := 1 ;i<n;i++{
		temp=cur 
		if s[i]=='0'{
			if s[i-1]!='1'&&s[i-1]!='2'{
				return 0
			}
			// 10 或者20 
			cur=pre
			
		}else if s[i-1]=='1'||(s[i-1]=='2'&&s[i]<='6'){
			fmt.Println(cur==pre,cur,pre)
			cur=cur+pre
		}
		pre=temp

	}
	return cur 
}




func main() {
	fmt.Println(numDecodings1("11111"))

}