package main

import (
	"fmt"
)


func countAndSay(n int) string {
	list := make([]string, n)
	list[0] = "1"
	for i := 1; i < n; i++ {
		list[i] = GetStrbyLast(list[i-1])
	}
	return list[n-1]

}

func GetStrbyLast(in string) string {
	out := []byte{}
	temp := in[0]
	num := 1
	l := len(in)
	for i := 1; i < l; i++ {
		if in[i] == temp {
			num++
		} else {
			out = append(out, byte(num+'0'))
			out = append(out, temp)
			temp = in[i]
			num = 1
		}
	}
	out = append(out, byte(num+'0'))
	out = append(out, temp)
	return string(out)
}

func main() {
	fmt.Println(countAndSay(6))

}
