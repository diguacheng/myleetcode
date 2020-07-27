package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i:=len((nums1))-1
	for m>0 && n>0{
		if nums1[m-1]>nums2[n-1]{
			nums1[i]=nums1[m-1]
			i--
			m--
		}else{
			nums1[i]=nums2[n-1]
			i--
			n--
		}
	}
	for n>0{
		nums1[i]=nums2[n-1]
		i--
		n--
	}
}
func main() {
	nums1:=[]int{1,2,3,0,0,0}
	m:= 3
	nums2:=[]int{2,5,6}
	n:= 3
	merge(nums1,m,nums2,n)
	fmt.Println(nums1)

}