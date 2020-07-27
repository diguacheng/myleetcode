package main

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max

}
func maxProfit1(prices []int) int {
	minpoint:=1<<31
	max := 0
	for i := 0; i < len(prices); i++ {
		if minpoint>prices[i] {
			minpoint=prices[i]
		}
		if prices[i]-minpoint>max{
			max=prices[i]-minpoint
		}
	}
	return max

}

func main() {

}
