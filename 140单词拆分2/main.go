package main

import (
	"fmt"
	"strings"
)

// 139题的题解 先检测是否可分 
func IFwordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 如果没有 139先用来检测可行性 会直接超时。 先检测可行性 再 用回溯算法求解 
func wordBreak(s string, wordDict []string) []string {
	if IFwordBreak(s,wordDict)==false{
		return []string{}
	}
	res := make([]string, 0)
	temp := make([]string, 0)
	table := make(map[string]bool)
	n := len(wordDict)
	for i := 0; i < n; i++ {
		table[wordDict[i]] = true
	}
	backTrack(s, table, temp, &res)
	return res
}

func backTrack(s string, table map[string]bool, temp []string, res *[]string) {
	n := len(s)
	if n == 0 {
		ans := make([]string, len(temp))
		copy(ans, temp)
		*res = append(*res, strings.Join(ans, " "))
		return
	}
	for i := 1; i <= n; i++ {
		if table[s[:i]] == true {
			temp = append(temp, s[:i])
			backTrack(s[i:], table, temp, res)
			temp = temp[:len(temp)-1]
		}
	}
	return
}

func wordBreak1(s string, wordDict []string) []string {
	res := make([]string, 0)
	temp := make([]string, 0)
	backTrack2(s, wordDict, temp, &res)
	return res
}
func backTrack2(s string, wordDict []string, temp []string, res *[]string) {
	n := len(s)
	if n == 0 {
		ans := make([]string, len(temp))
		copy(ans, temp)
		*res = append(*res, strings.Join(ans, " "))
		return
	}
	m := len(wordDict)
	var wordLen int
	for i := 0; i < m; i++ {
		wordLen = len(wordDict[i])
		if wordLen <= n && s[:wordLen] == wordDict[i] {
			temp = append(temp, wordDict[i])
			backTrack2(s[wordLen:], wordDict, temp, res)
			temp = temp[:len(temp)-1]
		}
	}
	return
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak1(s, wordDict))
}
