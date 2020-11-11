package main

import "fmt"

func findRotateSteps(ring string, key string) int {
	idxTable := map[byte][]int{}
	n := len(ring)
	counts := make([][]int, n)
	l := len(key)
	for i := 0; i < n; i++ {
		if _, ok := idxTable[ring[i]]; !ok {
			idxTable[ring[i]] = make([]int,0 )
		}
		idxTable[ring[i]] = append(idxTable[ring[i]], i)
		counts[i] = make([]int,l)
	}

	var dp func(start, i int) int
	dp = func(start, i int) int {
		if i >= l {
			return 0
		}
		if counts[start][i] > 0 {
			return counts[start][i]
		}
		ends:= idxTable[key[i]]
		min := 100000
		for _, end := range ends {
			mintemp := minL(start, end, n)+1+dp(end, i+1)
			if min > mintemp{
				min = mintemp
			}
		}
		counts[start][i] = min
		return min
	}

	return dp(0, 0)
}

func minL(a, b, n int) int {
	x := (a - b + n) % n
	y := (b - a + n) % n
	if x < y {
		return x
	}
	return y
}

func main() {
	ring := "godding"
	key := "gd"
	fmt.Println(findRotateSteps(ring,key))

}