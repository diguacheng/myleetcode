package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}
	revertedNumber := 0 //反转一半的数
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10

}
func main() {
	fmt.Println(isPalindrome(12431))

}
