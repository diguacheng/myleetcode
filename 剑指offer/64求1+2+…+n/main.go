package main

func sumNums(n int) int {
	// 递归
	ans := 0
	var f func(int) bool
	f = func(i int) bool {
		ans += n
		return n > 0 && f(n-1)
	}
	f(n)
	return ans
}

func main() {

}
