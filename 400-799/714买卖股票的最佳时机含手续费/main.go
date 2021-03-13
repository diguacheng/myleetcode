package main

//  动态规划
// 定义状态
// dp[i][0] 第i天交易后 手里没股票的最大收益
// dp[i][1] 第i天交易后 手里有股票的最大收益
// 则dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i]-fee)
//   dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
// 初始状态 为 dp[0][0]=0 dp[0][1]=-prices[0]
func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]

}

// 动态方程的优化
func maxProfit2(prices []int, fee int) int {
	n := len(prices)
	sell, buy := 0, -prices[0]

	for i := 1; i < n; i++ {
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, sell-prices[i])
	}
	return sell

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心算法
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
