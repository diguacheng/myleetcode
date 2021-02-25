package main

import "fmt"

//动态规划
// dp[i]代表字符串s[0..i]是否可以被空格拆分为一个或多个在字典中出现的单词
// 枚举 s[0..i-1]s[0..i−1] 中的分割点 jj ，看 s[0..j-1]s[0..j−1] 组成的字符串 s_1 s_2 是否都合法
//转移公式为 dp[i]=dp[j]&&isWod(s[j...i-1])
// isWod(s[j...i-1])检查s[j...i-1]是否在字典中

func wordBreak(s string, wordDict []string) bool {
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

func wordBreak1(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	n := len(s)
	list := make([]int, n)
	var help func(i int) bool
	help = func(i int) bool {
		if i == -1 {
			return true
		}
		if list[i] == 1 {
			return true
		}
		if list[i] == -1 {
			return false
		}

		res := false
		for k := i; k >= 0; k-- {
			if wordDictSet[s[k:i+1]] == true {
				res = help(k-1) || res
			}
		}
		if res == true {
			list[i] = 1
		} else {
			list[i] = -1
		}
		return res
	}
	return help(n - 1)

}

func main() {
	// s := "catsandog"
	// wordDict := []string{"cats", "dog","sand","and","cat"}
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
	fmt.Println(wordBreak1(s, wordDict))

}
