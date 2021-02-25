package main

func getRow(rowIndex int) []int {
	numRows := rowIndex
	if numRows == 0 {
		return []int{}
	}
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
	}
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		res[i][0] = 1
		res[i][i+1-1] = 1
		for j := 1; j < (i+2)/2; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
			res[i][i-j] = res[i][j]
		}
	}
	return res[rowIndex-1]
}

func main() {

}
