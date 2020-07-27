package main
func productExceptSelf(nums []int) []int {
	length := len(nums)
	l:=make([]int, length)
	r:=make([]int, length)
	ans:=make([]int, length)
	l[0]=1
	for i:=1;i<length;i++ {
		l[i]=l[i-1]*nums[i-1]
	}
	r[length-1]=1
	for i:=length-2;i>=0;i--{
		r[i]=r[i+1]*nums[i+1]
	}
	for i:=0;i<length;i++{
		ans[i]=l[i]*r[i]
	}
	return ans
}

func productExceptSelf1(nums []int) []int {
	length := len(nums)
	
	ans:=make([]int, length)
	ans[0]=1
	for i:=1;i<length;i++ {
		ans[i]=ans[i-1]*nums[i-1]
	}
	r:=1
	for i:=length-1;i>=0;i--{
		ans[i]=ans[i]*r
		r=r*nums[i]
	}
	return ans
}

func main(){

}