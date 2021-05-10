package main

import "fmt"

func minimumTimeRequired(jobs []int, k int) int {
	ans := 1<<31
	sum := make([]int, k)
	DFS(0, k, 0, 0,sum, jobs, &ans)
	return ans
}

func DFS(u, k, maxv ,uesd int, sum, jobs []int, as *int) {
	n:=len(jobs)
	ans := *as
	if maxv >= ans {
		return
	}
	if u == n {
		*as = maxv
		return
	}
	if uesd<k{
		sum[uesd] = jobs[u]
		DFS(u+1, k, max(maxv, sum[uesd]),uesd+1, sum, jobs, as)
		sum[uesd] =0
	}
	for i := 0; i < uesd; i++ {
		sum[i] += jobs[u]
		DFS(u+1, k, max(maxv, sum[i]),uesd, sum, jobs, as)
		sum[i] -= jobs[u]
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	arr := []int{1,2,4,7,8}
	k := 2
	ans := minimumTimeRequired(arr, k)
	fmt.Println(ans )

}