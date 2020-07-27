package main

func maxProfit(prices []int) int {
	max:=0
	min:=1<<31
	n:=len(prices)
	for i:=0;i<n;i++{
		if min>prices[i]{
			min=prices[i]
		}
		if prices[i]-min>max{
			max=prices[i]-min
		}
	}
	return max

}
func main() {

}
