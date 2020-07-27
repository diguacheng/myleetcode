package main

import "fmt"

// 和为s的两个数字

func twoSum(nums []int, target int) []int {
	start:=0
	end:=len(nums)-1
	for start<end{
		if nums[start]+nums[end]<target{
			start++
		}else if nums[start]+nums[end]>target{
			end--
			
		}else{
			break
		}

	}
	return []int{nums[start],nums[end]}

}

func findContinuousSequence(target int) [][]int {
	if target==1||target==2{
		return [][]int{{}}
	}
	res:=[][]int{}
	small,big :=1,2
	mid:=(1+target)/2
	curSum:=small+big
	for small<mid{
		if curSum==target{
			temp:=[]int{}
			for i:=small;i<=big;i++{
				temp=append(temp,i)
			}
			res=append(res,temp)
		}
		for curSum>target&&small<mid{
			curSum-=small
			small++
			if curSum==target{
				temp:=[]int{}
				for i:=small;i<=big;i++{
					temp=append(temp,i)
				}
				res=append(res,temp)

			}
			
		}
		big++
		curSum+=big
	}
	return res

}

func main() {
	//nums :=[]int {2,7,11,15}
	target := 9
	//fmt.Println(twoSum(nums,target))
	fmt.Println(findContinuousSequence(target))

}