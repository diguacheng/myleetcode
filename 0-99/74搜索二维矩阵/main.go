package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	start, end := 0, m-1
	var mid int
	if m!=1{
		for start <= end {
			mid = (start + end) / 2
			if matrix[mid][0] == target {
				return true
			}
			if matrix[mid][0] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	
	}
	i := end
	if i<0{
		return false
	}
	start, end = 0, n-1
	if n!=1{
		for start <= end {
			mid = (start + end) / 2
			if matrix[i][mid] == target {
				return true
			}
			if matrix[i][mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		return false
	}
	return matrix[i][0]== target
	
	
}


func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])

	start, end := 0, m*n-1
	var mid int
	for start <= end {
		mid=start + (end-start)/2

		if matrix[mid/n][mid-(mid/n)*n]==target{
			return true
		}
		if matrix[mid/n][mid-(mid/n)*n]>target{
			end=mid-1
		}else{
			start=mid+1
		}
	}
	return false
	
	
}
func main() {
	// matrix := [][]int{
	// 	{1, 3, 5, 7},
	// 	{10, 11, 16, 20},
	// 	{23, 30, 34, 50},
	// }
	matrix:=[][]int{
		{1},{3},
	}	
	fmt.Println(searchMatrix1(matrix, 0))
	fmt.Println(searchMatrix1(matrix, 21))

}
