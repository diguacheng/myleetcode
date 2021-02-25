package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
	}
	if sum%3 != 0 {
		return false
	}
	aver := sum / 3
	i := 0
	j := len(A) - 1
	left, right := A[i], A[j]
	for left != aver && i < len(A)-1 {
		left = left + A[i+1]
		i = i + 1
	}
	for right != aver && j > 0 {
		right = right + A[j-1]
		j = j - 1
	}
	if right == aver && left == aver && i+1 < j {
		return true
	}
	return false

}
func canThreePartsEqualSum1(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
	}
	if sum%3 != 0 {
		return false
	}
	aver := sum / 3
	i := 0
	cur := 0
	for i < len(A) {
		cur += A[i]
		if cur == aver {
			break
		}
		i++
	}
	if cur != aver {
		return false
	}
	j := i + 1
	for j < len(A)-1 {
		cur = cur + A[j]
		if cur == aver*2 {
			return true
		}
		j++
	}
	return false
}

func main() {
	A := []int{1, -1, 1, -1}
	fmt.Println(canThreePartsEqualSum(A))

}
