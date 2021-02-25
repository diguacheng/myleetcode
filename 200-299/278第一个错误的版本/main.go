package main

func firstBadVersion(n int) int {
	start, end := 1, n
	for start < end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}

	}
	return start
}
func main() {

}
