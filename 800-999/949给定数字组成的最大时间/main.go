package main

import (
	"fmt"
	"strconv"
)

func largestTimeFromDigits(arr []int) string {

	var h, m int
	maxH := -1
	maxM := -1
	for i := 0; i < 4; i++ {
		if arr[i] >= 3 {
			continue
		}
		for j := 0; j < 4; j++ {

			if i == j || arr[i]*10+arr[j] >= 24 {
				continue
			}
			for k := 0; k < 4; k++ {
				if k == i || k == j || arr[k] >= 6 {
					continue
				}
				for l := 0; l < 4; l++ {
					if l == i || l == j || l == k || arr[k]*10+arr[l] >= 60 {
						continue
					}

					h = arr[i]*10 + arr[j]
					m = arr[k]*10 + arr[l]
					if h > maxH || (h == maxH && m > maxM) {
						maxH = h
						maxM = m
					}
				}
			}
		}
	}
	if maxH == -1 {
		return ""
	}
	hs := strconv.Itoa(maxH)
	ws := strconv.Itoa(maxM)
	if len(hs) == 1 {
		hs = "0" + hs
	}
	if len(ws) == 1 {
		ws = "0" + ws
	}
	return hs + ":" + ws

}

func main() {
	arr := []int{0, 0, 1, 0}
	fmt.Println(largestTimeFromDigits(arr))
}
