package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	temp := ""
	help(temp, &res, 0, 0, n)
	return res

}

func help(temp string, res *[]string, in, out, n int) {
	if len(temp) == 2*n {

		*res = append(*res, temp)
		return
	}
	if in < n {
		temp = temp + "("
		help(temp, res, in+1, out, n)
		temp = temp[:len(temp)-1]
	}
	if out < in {
		temp = temp + ")"
		help(temp, res, in, out+1, n)
		temp = temp[:len(temp)-1]
	}

}

func main() {
	fmt.Println(generateParenthesis(3))

}
