package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	n := len(name)
	m := len(typed)
	if n == 0 || m == 0 {
		return false
	}
	if name[0] != typed[0] {
		return false
	}
	i, j := 1, 1
	for j < m {
		if i < n && name[i] == typed[j] {
			i++
			j++
		} else if typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == n

}

func main() {
	fmt.Println(isLongPressedName("alex", "aaleelx"))

}
