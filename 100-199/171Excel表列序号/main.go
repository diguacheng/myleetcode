package main

import "fmt"

func titleToNumber(s string) int {
	n := len(s)
	sum := 0
	c := 1
	for i := n - 1; i >= 0; i-- {
		sum += int(s[i]-'A'+1) * c
		c = c * 26
	}
	return sum
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))

}
