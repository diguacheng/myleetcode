package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	// 超时了
	ans := 0
	pre := make([]int, len(A)+1)
	pre[0] = 0 // pre[i]表示第i-1个元素的前缀和
	for i := 1; i <= len(A); i++ {
		pre[i] = pre[i-1] + A[i-1]
		for j := i - 1; j >= 0; j-- {
			if (pre[i]-pre[j])%K == 0 {
				ans++
			}
		}
	}
	return ans

}

func subarraysDivByK1(A []int, K int) int {
	// 前缀和模k位哈希表的k ，个数为value
	// p[i]表示钱i个元素的和
	// (p[i]-p[j])%k==0等价于p[i]%k==p[j]%k。
	// 也就是说 如果有两个数的余数相同，这个区间就满足和为k的倍数。
	// 初始余数为0对应为1，因为余数为0 本身可以是个子数组
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum%K + K) % K
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}

func subarraysDivByK2(A []int, K int) int {

	record := map[int]int{0: 1}
	sum, ans := 0, 0
	var modulus int
	for _, elem := range A {
		sum += elem
		modulus = (sum%K + K) % K
		record[modulus]++
	}
	for _, v := range record {
		ans += v * (v - 1) / 2
	}
	return ans
}

func main() {
	A := []int{4, 5, 0, -2, -3, 1}
	K := 5
	fmt.Println(subarraysDivByK1(A, K))
	fmt.Println(-3 % 5)

}
