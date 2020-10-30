package main

import (
	"fmt"
	"strconv")


func multiply(num1 string, num2 string) string {
	l1 := len(num1)-1
	l2 := len(num2)-1
	res := 0
	for i := l1 ; i >= 0; i-- {
		for j := l2  ; j >= 0; j-- {
			res += int(num1[i]-'0') * pow(10,l1-i) * int(num2[j]-'0') * pow(10,l2-j)
		}
	}
	return strconv.Itoa(res)
}

func pow(x, n int)int{
	res:=1
	for n>0{
		res*=10
		n--
	}
	return res
}



func main() {
	fmt.Println(multiply("123","4567"))

}