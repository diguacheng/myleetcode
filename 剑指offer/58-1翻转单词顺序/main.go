package main

func reverseWords(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	curr := make([]byte, 0)
	tmp := make([]string, 0)
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			if len(curr) > 0 {
				tmp = append(tmp, string(curr))
				curr = make([]byte, 0)
			}
		} else {
			curr = append(curr, s[i])
		}
	}
	if len(curr) > 0 {
		tmp = append(tmp, string(curr))
		curr = make([]byte, 0)
	}
	n = len(tmp)
	for i := 0; i < n/2; i++ {
		tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
	}
	res := ""
	for i := 0; i < n; i++ {
		res = res + tmp[i] + " "
	}
	if len(res) == 0 {
		return ""
	}
	return res[:len(res)-1]

}
func main() {

}
