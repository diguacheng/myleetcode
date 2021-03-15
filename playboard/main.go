package main

import "fmt"

func isN(num int) bool {
	if num == 1 || num == 2 {
		return true
	}
	if num%2 == 0 {
		return isN(num / 2)
	}
	return false
}

func main() {
	n:=1025

	for i := 1; i < n; i++ {
		if isN()
	}

}
