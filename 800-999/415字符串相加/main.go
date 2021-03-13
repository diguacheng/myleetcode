package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	carry := 0
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	l := l1
	if l1 < l2 {
		l = l2
	}
	res := make([]byte, l+2)
	var temp int
	for l1 >= 0 && l2 >= 0 {
		temp = int(num1[l1]-'0'+num2[l2]-'0') + carry
		carry = temp / 10
		temp = temp % 10
		res[l+1] = byte(temp + '0')
		l--
		l1--
		l2--
	}
	for l1 >= 0 {
		temp = int(num1[l1]-'0') + carry
		carry = temp / 10
		temp = temp % 10
		res[l+1] = byte(temp + '0')
		l1--
		l--
	}
	for l2 >= 0 {
		temp = int(num2[l2]-'0') + carry
		carry = temp / 10
		temp = temp % 10
		res[l+1] = byte(temp + '0')
		l2--
		l--
	}
	if carry > 0 {
		res[l+1] = byte(carry + '0')
	}
	if res[0] == '0' || res[0] == 0 {
		res = res[1:]
	}
	return string(res)
}
func main() {
	num1 := "6"
	num2 := "501"
	fmt.Println(addStrings(num1, num2))

}
