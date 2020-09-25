package main

import "fmt"

func zeroOnePack(v,w []int,m int)int{
	n:=len(v)
	dp:=make([][]int,n+1)
	for i:=0; i<=n; i++{
		dp[i]=make([]int,m+1)
	}
	for i:=1;i<=n;i++{
		for j:=0; j <=m;j++{
			dp[i][j] =dp[i-1][j]
			if (j>=v[i-1]){
				dp[i][j]=max(dp[i][j],dp[i-1][j-v[i-1]]+w[i-1])
			}
		}
	}
	return dp[n][m]
}

func zeroOnePack1(v,w []int,m int)int{
	//用1维数组 
	n:=len(v)
	dp:=make([]int,m+1)
	for i:=1;i<=n;i++{
		for j:=m; j>=v[i-1];j--{
			dp[j]=max(dp[j],dp[j-v[i-1]]+w[i-1])
		}
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[i][j]表示只有前i个物品，总体积是j的情况下，总价值最大的值
// result =max{dp[n][0-v]}
// 1. 不选第i个物品 dp[i][j]=dp[i-1][j]
// 2 选第i个物品 dp[i][j]=dp[i-1][j-v[i]]+w[i]
// dp[i][j]=max(dp[i-1][j],dp[i-1][j-v[i]]+w[i])
// dp[0][0]=0

func main() {
	// var n, m int
	// var v [1010]int
	// var w [1010]int
	// fmt.Scanln(&n, &m)
	// //fmt.Println(n,m)
	// for i := 1; i <= n; i++ {
	// 	fmt.Scanln(&v[i], &w[i])
	// }
	// fmt.Println(v[:n],w[:n])

	// dp := make([][]int, n+1)
	// for i := 0; i <= n; i++ {
	// 	dp[i] = make([]int, m+1)
	// }
	// for i := 1; i <= n; i++ {
	// 	for j := 0; j <= m; j++ {
	// 		dp[i][j] = dp[i-1][j]
	// 		if j >= v[i] {
	// 			dp[i][j] = max(dp[i][j], dp[i-1][j-v[i]]+w[i])
	// 		}
	// 	}
	// }
	// res := 0
	// for i := 0; i <= m; i++ {
	// 	res = max(res, dp[n][i])
	// }
	// fmt.Println(res)

	v := []int{1, 2, 3, 4}
	w := []int{2, 4, 4, 5}
	fmt.Println(zeroOnePack(v, w, 5))
	fmt.Println(zeroOnePack1(v, w, 5))
}
