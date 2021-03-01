package main

type NumArray struct {
    arr []int
}


func Constructor(nums []int) NumArray {
    if len(nums)==0{
        return NumArray{
            arr:make([]int,0),
        }
    }
    arr:=make([]int,len(nums))

    arr[0]=nums[0]
    for i:=1;i<len(nums);i++{
        arr[i]=arr[i-1]+nums[i]
    }
    return NumArray{
        arr:arr,
    }
}


func (this *NumArray) SumRange(i int, j int) int {
    if i==0{
        return this.arr[j]
    }
    return this.arr[j]-this.arr[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */