package main

import "fmt"

func fraction(cont []int) []int {
	n := len(cont)
	res := []int{cont[n-1], 1}
	for i := n - 2; i >= 0; i-- {
		temp := res[1]
		res[1] = res[0]
		res[0] = cont[i]*res[0] + temp
		fmt.Println(res)

	}
	return res
}
func main() {
	cont := []int{3, 2, 0, 2}
	res := fraction(cont)
	fmt.Println(res)
}
