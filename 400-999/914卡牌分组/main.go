package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	// 暴力法
	l := len(deck)
	for x := 2; x <= len(deck); x++ {
		if l%x != 0 {
			continue
		}
		mmap := make(map[int]int)
		for _, v := range deck {
			mmap[v]++
		}
		count := 0
		for _, v := range mmap {
			if v%x == 0 {
				count++
			}
		}
		if count == len(mmap) {
			return true
		}
	}
	return false

}

func hasGroupsSizeX1(deck []int) bool {
	mmap := make(map[int]int)
	for _, v := range deck {
		mmap[v]++
	}
	g := -1
	for _, v := range mmap {
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g>=2
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func main() {
	a := []int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(hasGroupsSizeX1(a))

}
