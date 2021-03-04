package main

import (
	"fmt"
	"sort")


func maxEnvelopes(envelopes [][]int) int {
	m := len(envelopes)
	if m == 1 {
		return 1
	}
	sort.Slice(envelopes, func(i, j int) bool {
		w1, h1 := envelopes[i][0], envelopes[i][1]
		w2, h2 := envelopes[j][0], envelopes[j][1]
		return w1*h1 < w2*h2
	})
	dp := make([]int, m)
	dp[0] = 1
	for i := 1; i < m; i++ {
	
		for j := 0; j < i; j++ {
			w1, h1 := envelopes[i][0], envelopes[i][1]
			w2, h2 := envelopes[j][0], envelopes[j][1]
			if w1> w2 && h1 > h2 && dp[i]<= dp[j] {
				dp[i]=dp[j]+1
			}
		}
		
	}
	res := 1
	for i := 0; i < m; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func main() {
	e:=[][]int{{5,4},{6,4},{6,7},{2,3}}
	fmt.Println(maxEnvelopes(e))


}