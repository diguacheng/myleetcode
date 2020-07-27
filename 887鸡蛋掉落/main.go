package main

import "fmt"

var mmap map[int]int

func superEggDrop(K int, N int) int {
	mmap=make(map[int]int)
	return dp(K,N)
	

}

func dp(K,N int)int{
	if _,ok:=mmap[N*100+K];!ok{
		ans:=0
		if N==0{
			ans=0
		}else if K==1{
			ans=N
		}else{
			low,high:=1,N
			for low+1<high{
				x:=(low+high)/2
				t1:=dp(K-1,x-1)
				t2:=dp(K,N-x)
				if t1<t2{
					low=x
				}else if t1>t2{
					high=x
				}else{
					high=x
					low=x
				}
			}
			ans= 1+min(max(dp(K-1,low-1),dp(K,N-low)),max(dp(K-1,high-1),dp(K,N-high)))
		}
		mmap[N*100+K]=ans

	}
	return mmap[N*100+K]


}

func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}

func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}

func main() {
	fmt.Println(superEggDrop(1,2))
	fmt.Println(superEggDrop(2,6))
	fmt.Println(superEggDrop(3,14))
	

}
