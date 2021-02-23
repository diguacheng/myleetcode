package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	if n == 0 {
		return nil
	}
	res := make([]bool, n)
	maxx := 1
	for i := 0; i < n; i++ {
		if candies[i] > maxx {
			maxx = candies[i]
		}
	}
	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= maxx {
			res[i] = true

		}
	}
	return res

}

func mian() {

}
