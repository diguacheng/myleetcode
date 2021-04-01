package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 冒泡排序
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

// 选择排序
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// 插入排序
func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		current := arr[i]
		index := i - 1
		for index >= 0 && arr[index] > current {
			arr[index+1] = arr[index]
			index--
		}
		arr[index+1] = current
	}
	return arr
}

// 希尔排序
func shellSort(arr []int) []int {
	n := len(arr)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			j := i
			current := arr[i]
			for j-gap >= 0 && current < arr[j-gap] {
				arr[j] = arr[j-gap]
				j = j - gap
			}
			arr[j] = current
		}
	}
	return arr
}

// 归并排序
func mergeSort(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	mid := len / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	res := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	for len(left) > 0 {
		res = append(res, left...)
		left = nil
	}
	for len(right) > 0 {
		res = append(res, right...)
		right = nil
	}
	return res
}

// 快速排序
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, start, end int) {
	if start < end {
		pivot := partition(arr, start, end)
		quickSort(arr, start, pivot-1)
		quickSort(arr, pivot+1, end)
	}
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	for start < end {
		for start < end && arr[start] <= pivot {
			start++
		}
		arr[end] = arr[start]
		for start < end && arr[end] >= pivot {
			end--
		}
		arr[start] = arr[end]
	}
	arr[end] = pivot
	return end
}

// 堆排序
func heapSort(arr []int) []int {
	n := len(arr)
	heapify := func(arr []int, k int) {
		for {
			left := 2*k + 1  // 左子树下标
			right := 2*k + 2 // 右子树下标
			idx := k         // 根 左 右 最大值的索引
			if left < n && arr[left] > arr[idx] {
				idx = left
			}
			if right < n && arr[right] > arr[idx] {
				idx = right
			}
			if idx == k {
				// 若根节点较大，不用将节点下沉
				break
			}
			arr[k], arr[idx] = arr[idx], arr[k]
			// 交换过后，idx保存的不是最大值的索引，是从根节点换下的索引
			k = idx
			// 继续节点下沉
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}
	for i := n - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		n-- // 注意这里的长度要减去一
		heapify(arr, 0)
	}
	return arr
}

// 数组没有重读的数 且最后是有具体范围的数
// 计数排序
func countingSort(arr []int) []int {
	array := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
		array[count] = arr[i]
	}
	return array
}

// 数组有重复的数
func countingSort1(arr []int) []int {
	minValue, maxValue := 1<<31, -1<<31
	for i := 0; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
		}
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	bucket := make([]int, maxValue-minValue+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]-minValue]++
	}
	i := 0
	for j := 0; j < len(bucket); j++ {
		for bucket[j] > 0 {
			arr[i] = j + minValue
			i++
			bucket[j]--
		}

	}
	return arr
}

// 桶排序
func bucktSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	n := len(arr)
	var i int
	// 找出待排序数组的最大值，最小值，一般这个是已知的
	minValue := arr[0]
	maxValue := arr[1]
	for i = 1; i < n; i++ {
		if arr[i] < minValue {
			minValue = arr[i]
		}
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}

	bucketSize := 100 // 桶的数量，可以子集定义
	bucketcount := (maxValue-minValue)/bucketSize + 1
	buckets := make([][]int, bucketcount)
	for i := 0; i < bucketcount; i++ {
		buckets[i] = make([]int, 0)
	}

	for i := 0; i < len(arr); i++ {
		// 将元素放入到对应得桶里面
		buckets[(arr[i]-minValue)/bucketSize] = append(buckets[(arr[i]-minValue)/bucketSize], arr[i])
	}
	res := make([]int, 0, len(arr))
	for i := 0; i < len(buckets); i++ {
		res = append(res, QuickSort(buckets[i])...)
	}
	return res

}

func radixSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	radix := 10 // 基数
	maxValue := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	// 求出最大值的长度决定
	maxDigit := len(strconv.FormatInt(int64(maxValue), 10))
	divisor := 1 // 定义每一轮的除数 1，10，100.。。

	// counter:=make([]int,10) // 保存每个桶中的元素个数
	for i := 0; i < maxDigit; i++ {
		bucket := make([][]int, radix) // 注意 没轮都要清空

		for j := 0; j < n; j++ {
			digit := (arr[j] / divisor) % radix
			bucket[digit] = append(bucket[digit], arr[j])

		}
		arr = []int{}
		for _, v := range bucket {
			arr = append(arr, v...)
		}

		divisor *= 10
	}
	return arr
}

func exectime(s func([]int) []int, in []int) {
	start := time.Now()
	_ = s(in)
	dur := time.Since(start)
	//fmt.Println(res)
	fmt.Println(dur)

}

// 生成a-b 得不重复得序列
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func t() {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}
	for i := 19; i >= 1; i-- {
		temp := a[rand.Intn(math.MaxInt64)%20]
		a[rand.Intn(math.MaxInt64)%20] = a[i]
		a[i] = temp

	}
	fmt.Println(a)
}
func main() {
	arr := generateRandomNumber(1, 11, 10)
	exectime(shellSort, arr)
	exectime(radixSort, arr)
	exectime(QuickSort, arr)

}
