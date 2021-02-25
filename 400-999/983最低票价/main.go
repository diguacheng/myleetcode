package main

import (
	"fmt"
	"math"
)

func mincostTickets(days []int, costs []int) int {
	memory := [366]int{}
	daym := map[int]bool{}
	for _, d := range days {
		// 将需要旅游的天数标记为true
		daym[d] = true
	}
	var dp func(day int) int
	// dp(day)表示从第day天到一年的结束 需要花费的钱。
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if memory[day] > 0 {
			return memory[day]
		}
		if daym[day] {
			memory[day] = min(min(dp(day+1)+costs[0], dp(day+7)+costs[1]), dp(day+30)+costs[2])
		} else {
			// 如果这一天不是必须出行，就不出行，就不买
			memory[day] = dp(day + 1)
		}
		return memory[day]
	}
	return dp(1)

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func mincostTickets1(days []int, costs []int) int {
	memory := [366]int{}
	durations := []int{1, 7, 30}
	var dp func(i int) int
	dp = func(i int) int {
		if i >= len(days) {
			return 0
		}
		if memory[i] > 0 {
			return memory[i]
		}
		memory[i] = math.MaxInt32
		j := i
		for k := 0; k < 3; k++ {
			for ; j < len(days) && days[j] < days[i]+durations[k]; j++ {
			}
			memory[i] = min(memory[i], dp(j)+costs[k])
		}
		return memory[i]

	}
	return dp(0)
}

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	fmt.Println(mincostTickets1(days, costs))
}
