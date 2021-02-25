package mian

func guessNumber(n int) int {
	s, e := 1, n
	mid := (s + e) / 2
	x := guess(mid)
	for x != 0 {
		if x > 0 {
			s = mid + 1
		} else {
			e = mid - 1
		}
		mid = (s + e) / 2
		x = guess(mid)

	}
	return mid

}

func main() {

}
