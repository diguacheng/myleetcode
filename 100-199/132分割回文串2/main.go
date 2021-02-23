package main

import "fmt"

func minCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	//table:=make(map[string]int)
	check := func(s string) bool {
		n := len(s)
		if n == 1 {
			return true
		}
		l := n / 2
		start,end:=l-1,l
		if n%2==1{
			end++
		}
		for start>=0{
			if s[start]!=s[end]{
				return false
			}
			start-- 
			end++
		}

		return true
	}
	dp := make([]int, n+1)
	dp[0] = -1
	dp[1] = 0
	var min int
	for i := 2; i <=n; i++ {
		min = n + 1
		for j := 1; j < i; j++ {
			if check(s[j-1:i]) {
				if min > dp[j-1]+1 {
					min = dp[j-1] + 1
				}
			}
		}
		if min > dp[i-1]+1 {
			min = dp[i-1] + 1
		}
		dp[i] = min
	}
	return dp[n]
}

func main() {
	s := "fifgbeajcacehiicccfecbfhhgfiiecdcjjffbghdidbhbdbfbfjccgbbdcjheccfbhafehieabbdfeigbiaggchaeghaijfbjhi"
	fmt.Println(minCut(s))

}