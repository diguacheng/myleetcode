package main

import "fmt"

func removeDuplicates(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}
	st := make([]byte, 1)
	st[0] = S[0]
	i := 1
	for i < n {
		if len(st) == 0 || st[len(st)-1] != S[i] {
			st = append(st, S[i])
		} else if len(st) > 0 && st[len(st)-1] == S[i] {
			st = st[:len(st)-1]
		}
		i++
	}
	return string(st)
}

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates(s))
}
