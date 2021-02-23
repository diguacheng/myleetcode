package main

import "fmt"

func change(amount int, coins []int) int {
	// 超时 
	n := len(coins)
	res := 0
	var help func(i int, temp int)
	help = func(i int, temp int) {

		if temp == amount {
			res++
			return
		}
		for k := i; k < n; k++ {
			if temp+coins[k]<=amount {
				help(k, temp+coins[k])
			}
			
		}
	}
	help(0, 0)
	return res

}

func change1(amount int, coins []int) int {
	// dp 
	n := len(coins)
	dp:=make([]int,amount+1)
	dp[0]=1
	for i := 0; i < n; i++ {
		for x:=coins[i];x<=amount;x++{
			dp[x]+=dp[x-coins[i]]
		}
	}
	return dp[amount]

}


func main() {
	fmt.Println(change(500,[]int{3,5,7,8,9,10,11}))
	fmt.Println(change1(500,[]int{3,5,7,8,9,10,11}))

}