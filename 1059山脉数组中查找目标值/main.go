package main

import "fmt"

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
	array []int

}

func Creater(nums []int)MountainArray{
	return MountainArray{nums,}
}

func (this *MountainArray) get(index int) int {
	return this.array[index]
}
func (this *MountainArray) length() int       {
	return len(this.array)
}

func (this *MountainArray) getmaxindex() int {

	n := this.length() - 1
	div := n / 10+1
	maxindex := 0
	for i := 1; i <= n; i += div {
		a := this.get(i)
		if b := this.get(i - 1); a < b {
			i = i - div
			div = div / 10
			if div == 0 {
				div = 1
			}
			a = b
		}
		if div == 1 {
			i = i + 1
			for i <= n &&this.get(i) > this.get(i-1) {
				i++
			}
			if this.get(i)>this.get(i-1){
				maxindex=i
			}else{
				maxindex =i-1
			}

			break
		}
	}
	return maxindex

}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	maxindex := mountainArr.getmaxindex()
	n := mountainArr.length() - 1
	start, end := 0, maxindex
	a, b := -1,-1
	for start <= end {
		mid := (start + end) / 2
		temp := mountainArr.get(mid)
		if temp == target {
			a = mid
			break
		}
		if temp > target {
			end = mid-1
		} else {
			start = mid+1

		}
	}
	start, end = maxindex, n
	for start <= end {
		mid := (start + end) / 2
		temp := mountainArr.get(mid)
		if temp == target {
			b = mid
			break
		}
		if temp < target {
			end = mid-1
		} else {
			start = mid+1
		}
	}
	if a >= 0 {
		return a
	}
	return b

}

func main() {
	array :=[]int{1,2,3,4,5,3,1} 
	target := 2
	m:=Creater(array)
	fmt.Println(m.getmaxindex())
	fmt.Println(findInMountainArray(target,&m))

}
