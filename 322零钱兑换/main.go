package main

import (
	"fmt"
)

// 给定不同面额的硬币 coins 和一个总金额 amount。
//编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
//如果没有任何一种硬币组合能组成总金额，返回 -1。

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := make([]int, amount+1)
	mincoin := coins[0]
	for _, v := range coins {
		if v > amount {
			continue
		}
		res[v] = 1
		if v < mincoin {
			mincoin = v
		}
	}
	if amount < mincoin {
		return -1
	}
	start := mincoin + 1
	for mincoin > 0 {
		mincoin--
		res[mincoin] = -1
	}
	for i := start; i <= amount; i++ {
		if res[i] > 0 {
			continue
		}
		min := 1 << 31
		for _, j := range coins {
			if j > i {
				continue
			}
			if res[i-j] != -1 && res[i-j]+1 <= min {
				min = res[i-j] + 1
			}
		}
		if min == 1<<31 {
			res[i] = -1
		} else {
			res[i] = min
		}

	}
	return res[amount]

}

func coinChange1(coins []int, amount int) int {
	// 自底向上 动态规划
	if amount == 0 {
		return 0
	}
	res := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		res[i] = amount + 1
		for _, v := range coins {
			if v <= i {
				if res[i] > res[i-v]+1 {
					res[i] = res[i-v] + 1
				}

			}
		}
	}
	if res[amount] > amount {
		return -1
	}
	return res[amount]

}

func coinChange2(coins []int, amount int) int {
	// 自上到下 动态规划
	if amount == 0 {
		return 0
	}
	res := make([]int, amount+1)
	return dp(coins, res, amount)

}

func dp(coins, res []int, amount int) int {
	// 这里的amount 是指 amount-某个硬币的值 如果小于0，说明 不应该使用这个硬币，需要返回-1，因为外面有个用掉这个硬币的加1
	// 如果等于0，说明这个amount 等于该硬币的值，不需要其他硬币，返回0，这个硬币已经在外面加1了
	//
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if res[amount] != 0 {
		// 如果已经在前面被计算过，就不用重新计算，直接返回结果就好。
		return res[amount]
	}
	min := 1 << 31
	for _, v := range coins {
		r := dp(coins, res, amount-v) + 1
		if r > 0 && r < min {
			min = r
		}

	}
	if min == 1<<31 {
		min = -1
	}
	res[amount] = min
	return min
}


func coinChange4(coins []int, amount int) int {
	// @data 20200922
	if amount == 0 {
		return 0
	}
	dp:= make([]int, amount+1)
	n:=len(coins)

	
	for i:=1 ;i<=amount;i++{
		dp[i]=amount+1
		for j:=0;j<n;j++{
			if coins[j]>i{
				continue
			}
			if dp[i-coins[j]]>=0{
				dp[i]=min(dp[i],1+dp[i-coins[j]])
			}
		}


	}
	if dp[amount]>amount{
		return-1
	}
	return dp[amount]
	
}



func min(a,b int) int{
	if a<b {
		return a
	}
	return b
}
func main() {
	coins := []int{474,83,404,3}
	amount := 264
	
	fmt.Println(coinChange4(coins, amount))
}
