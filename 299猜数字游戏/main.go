package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getHint(secret string, guess string) string {
	x, y := 0, 0
	table := make([]int, 10)
	n := len(secret)
	for i := 0; i < n; i++ {
		table[secret[i]-'0']++
		if guess[i] == secret[i] {
			table[guess[i]-'0']--
			x++
		}
	}
	for i := 0; i < n; i++ {
		if guess[i] != secret[i] {
			if table[guess[i]-'0'] > 0 {
				table[guess[i]-'0']--
				y++
			}
		}
	}
	return strings.Join([]string{strconv.Itoa(x), "A", strconv.Itoa(y), "B"}, "")

}

func main() {
	fmt.Println(getHint("1807", "7810"))

}
