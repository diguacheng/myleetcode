package main

func backspaceCompare(S string, T string) bool {
	return f(S) == f(T)
}

func f(s string) string {
	n := len(s)
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] != '#' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(stack)

}

func main() {

}
