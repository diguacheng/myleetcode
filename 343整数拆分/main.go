package main

import "fmt"

func integerBreak(n int) int {
	max := 0
	for i := 2; i <= n; i++ {
		c := counts(split(n, i), i)
		if max < c {
			max = c
		}
		if max > c {
			break
		}
	}

	return max
}

func split(n, i int) []int {
	res := make([]int, i)
	num := n / i
	r := n % i
	for j := 0; j < i; j++ {
		res[j] = num
		if r > 0 {
			res[j]++
			r--
		}
	}
	return res
}

func counts(in []int, i int) int {
	res := in[0]
	for j := 1; j < i; j++ {
		res = res * in[j]
	}
	return res

}


func integerBreak2(n int) int {
	if n<4{
		return n-1
	}
	t:=make([]int,n+1)
	t[0]=0
	t[1]=0
	var dp func(int)int
	dp=func(x int)int{
		if x==0||x==1{
			return 0
		}
		if t[x]!=0{
			if t[x]>x{
				return t[x]
			}
			return x
		}
		res:=0
		for i:=1;i<x;i++{
			temp:=dp(x-i)*dp(i)
			if x>temp{
				temp=x
			}
			if res<temp{
				res=temp
			}
		}
		t[x]=res
		return res
	}
	return dp(n)
}



func main() {
	fmt.Println(integerBreak2(20))

}


