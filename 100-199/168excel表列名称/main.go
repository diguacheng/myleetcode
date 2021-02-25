package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}

	ans := make([]byte, 0)

	for n > 0 {
		n--
		ans = append([]byte{byte(n%26 + 'A')}, ans...)
		n = n / 26

	}
	return string(ans)

}

func main() {
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))

}
