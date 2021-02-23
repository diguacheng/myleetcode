package main

import "fmt"

func getPermutation(n int, k int) string {
	candidates := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	c := candidates[:n]
	res := []string{}
	temp := make([]byte, 0)
	backTrack(c, temp, &res, k)
	//fmt.Println(res)
	return res[k-1]
}
func backTrack(c []byte, temp []byte, res *[]string, k int) {
	if len(*res) >= k {
		return
	}
	n := len(c)
	if n == 0 {
		ans := make([]byte, len(temp))
		copy(ans, temp)
		*res = append(*res, string(ans))
	}
	for i := 0; i < n; i++ {
		temp = append(temp, c[i])
		newc := append([]byte{}, c[0:i]...)
		newc = append(newc, c[i+1:]...)
		backTrack(newc, temp, res, k)
		temp = temp[:len(temp)-1]
	}
}

func getPermutation1(n int, k int) string {
	candidates := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	res := make([]byte, 0)
	c := candidates[:n]
	k--
	for n>0{
		nums := getFactorial(n - 1)
		a := k / nums
		k = k % nums
		var newc []byte
		res = append(res, c[a])
		newc = append([]byte{}, c[:a]...)
		newc = append(newc, c[a+1:]...)
		c = newc
		n--
	}

	return string(res)
}

func getFactorial(n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i

	}
	return res

}
func main() {
	fmt.Println(getPermutation1(3, 3))

}
