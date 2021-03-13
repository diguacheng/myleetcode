package main

import "fmt"

func longestMountain(A []int) int {
	n := len(A)
	maxl := 0
	i := 1
	for i < n {
		if A[i-1] < A[i] {
			j := i
			for j < n && A[j-1] < A[j] {
				j++
			}
			if j < n && A[j-1] == A[j] {
				i = j
				continue
			}
			for j < n && A[j-1] > A[j] {
				j++
			}
			if maxl > j-i {
				maxl = j - i
			}
			i = j
			continue
		}
		i++
	}
	return maxl

}

func main() {
	a := []int{2, 1, 4, 7, 3, 2, 5}
	fmt.Println(longestMountain(a))

}
