package main

import (
	"fmt"
	"time"
)

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] { // 相邻元素两两比较
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func BubbleSort1(arr []int) []int {
	n := len(arr)
	var flag bool
	for i := 0; i < n-1; i++ {
		flag = false
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] { // 相邻元素两两比较
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if flag == false {
			break
		}

	}
	return arr
}

func selectionSort(arr []int)[]int{
	n:=len(arr)
	for i:=0;i<n-1;i++{
		minIndex:=i
		for j:=i+1;j<n;j++{
			if arr[j]<arr[minIndex]{
				minIndex=j
			}
		}
		arr[i],arr[minIndex]=arr[minIndex],arr[i]
	}
	return arr
}

func insertionSort(arr []int)[]int{
	n:=len(arr)
	for i:=1;i<n;i++{
		current:=arr[i]
		index:=i-1
		for index>=0&& arr[index]>current{
			arr[index+1]=arr[index]
			index--
		}
		arr[index+1]=current
	}
	return arr
}

func shellSort(arr []int)[]int{
	n:=len(arr)
	for gap:=n/2;gap>0;gap=gap/2{
		for i:=gap;i<n;i++{
			j:=i
			current:=arr[i]
			for j-gap>=0&&current<arr[j-gap]{
				arr[j]=arr[j-gap]
				j=j-gap
			}
			arr[j]=current
		}
	}
	return arr 
}


func mergeSort(arr []int)[]int{
	len:=len(arr)
	if len<2{
		return arr 
	}
	mid:=len/2
	left:=arr[:mid]
	right:=arr[mid:]
	return merge(mergeSort(left),mergeSort(right))
}

func merge(left,right []int)[]int{
	res:=[]int{}
	for len(left)>0 && len(right)>0{
		if left[0]<=right[0]{
			res=append(res,left[0])
			left=left[1:]
		}else{
			res=append(res,right[0])
			right=right[1:]
		}
	}
	for len(left)>0{
		res=append(res,left...)
		left=nil
	}
	for len(right)>0{
		res=append(res,right...)
		right=nil
	}
	return res
}


func QuickSort(arr []int)[]int{
	quickSort(arr,0,len(arr)-1)
	return arr
}

func quickSort(arr []int,start,end int){
	if start<end{
		pivot:=partition(arr,start,end)
		quickSort(arr,start,pivot-1)
		quickSort(arr,pivot+1,end)
	}
}

func partition(arr []int,start,end int)int{
	pivot := arr[end]
	for start<end{
		for start < end &&arr[start]<=pivot{
			start++
		}
		arr[end]=arr[start]
		for start < end &&arr[end]>=pivot{
			end--
		}
		arr[start]=arr[end]
	}
	arr[end]=pivot
	return end
}

func heapSort(arr []int)[]int{
	n:=len(arr)
	heapify:=func(arr []int,k int){
		for {
			left:=2*k+1  // 左子树下标
			right:=2*k+2// 右子树下标
			idx:=k  // 根 左 右 最大值的索引
			if left<n&&arr[left]>arr[idx]{
				idx=left
			}
			if right<n&&arr[right]>arr[idx]{
				idx=right
			}
			if idx==k{
				// 若根节点较大，不用将节点下称
				break
			}
			arr[k],arr[idx]=arr[idx], arr[k]
			// 交换过后，idx保存的不是最大值的索引，是从根节点换下的索引
			k=idx
			// 继续节点下沉
		}
	}

	for i:=n/2-1;i>=0;i--{
		heapify(arr,i)
	}
	for i:=n-1;i>=1;i--{
		arr[0],arr[i]=arr[i], arr[0]
		n-- // 注意这里的长度要减去一 
		heapify(arr,0)
	}
	return arr
}

// 数组没有重读的数
func countingSort(arr []int)[]int{
	array:=make([]int,len(arr))
	for i := 0; i < len(arr); i++ {
		count:=0
		for j := 0; j < len(arr);j++ {
			if arr[i]>arr[j]{
				count++
			}
		}
		array[count]=arr[i]
	}
	return array
}

// 数组有重复的数
func countingSort1(arr []int)[]int{
	minValue,maxValue :=1<<31,-1<<31
	for i:=0;i<len(arr);i++{
		if arr[i]<minValue{
			minValue = arr[i]
		}
		if arr[i]>maxValue{
			maxValue = arr[i]
		}
	}
	bucket:=make([]int,maxValue-minValue+1)
	for i:=0;i<len(arr);i++{
		bucket[arr[i]-minValue]++
	}
	i:=0
	for j:=0;j<len(bucket);j++{
		for bucket[j]>0{
			arr[i]=j+minValue
			i++
			bucket[j]--
		}

	}
	return arr
}


func bucktSort(arr []int)[]int{
	if len(arr) == 0 {
		return arr
	}
	n:=len(arr)
	var i int
	minValue:=arr[0]
	maxValue:=arr[1]
	for i=1;i<n; i++ {
		if arr[i]<minValue{
			minValue=arr[i]
		}
		if arr[i]>maxValue{
			maxValue=arr[i]
		}
	}
	bucketSize:=5
	bucketcount:=(maxValue-minValue)/bucketSize+1
	buckets:=make([][]int,bucketcount)
	for i:=0; i < bucketcount; i++ {
		buckets[i] = make([]int,0)
	}

	for i:=0;i<len(arr);i++{
		buckets[(arr[i]-minValue)/bucketSize]=append(buckets[(arr[i]-minValue)/bucketSize], arr[i])
	}

	
	





}
func exectime(s func([]int) []int, in []int) {
	start := time.Now()
	res := s(in)
	dur := time.Since(start)
	fmt.Println(res)
	fmt.Println(dur)

}

func main() {
	arr := []int{8,6,4,9,9,4,7,5,3,3,2,1}
	exectime(countingSort1, arr)
	

}
