package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, 1)
	dp[0] = 1
	maxans := 1
	for i := 1; i < len(nums); i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if maxval < dp[j] {
					maxval = dp[j]

				}
			}
		}
		dp = append(dp, maxval+1)
		if maxans < dp[i] {
			maxans = dp[i]

		}
	}
	return maxans

}

func lengthOfLIS1(nums []int) int {
	n:=len(nums)
	if n<=1{
		return n
	}
	dp:=make([]int,1)
	dp[0]=1
	maxans:=1
	for i:=1;i<n;i++{
		maxVal:=0
		for j:=0;j<i;j++{
			if nums[i]>nums[j]&&maxVal<dp[j]{
				maxVal=dp[j]
			}
		}
		dp=append(dp, maxVal+1)
		if maxans<dp[i]{
			maxans=dp[i]
		}	
	}
	return maxans
}

func lengthOfLIS2(nums []int)int{
	n:=len(nums)
	if n<=1{
		return n
	}
	d:=make([]int,n+1)
	d[0]=nums[0]
	l:=1
	for i:=1;i<n;i++{
		if nums[i]>d[l]{
			l++
			d[l]=nums[i]
		}else{
			left,right:=1,l
			pos:=0
			for left<=right{
				mid:=left+(right-left)/2 
				if d[mid]<nums[i]{
					pos=mid
					left=mid+1
				}else{
					right=mid-1
				}
			}
			d[pos+1]=nums[i]
		}
	}
	return l

}

func main() {
	nums := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(lengthOfLIS1(nums))

}
