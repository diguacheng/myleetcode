package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	start, end := 0, len(arr)-1
	var mid int
	for start+1 < end {
		mid = (start + end) / 2
		if arr[mid] == x {
			break

		}
		if arr[mid] < x {
			start = mid
		} else {
			end = mid
		}
	}
	i, j := mid, mid+k-1
	if j > len(arr)-1 {
		i, j = i-(mid+k-len(arr)), j-(mid+k-len(arr))
	}
	for i > 0 && arr[j]-x >= x-arr[i-1] {
		j--
		i--
	}
	return arr[i : j+1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr,4,5))

}