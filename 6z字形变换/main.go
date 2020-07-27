package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	res := make([][]byte, numRows)
	k := 0
	for k < n {
		for i := 0; i < numRows && k < n; i++ {
			res[i] = append(res[i], s[k])
			k++
		}
		for i := numRows - 2; i > 0 && k < n; i-- {
			res[i] = append(res[i], s[k])
			k++
		}
	}
	ans :=[]byte{}
	for i := 0; i < numRows; i++ {
		ans =append(ans,res[i]...)
	}
	return string(ans)
}

func convert1(s string, numRows int) string {
	// 更慢了
	n := len(s)
	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]byte, n/2+1)
	}
	k := 0
	x := 0
	for k < n {
		for i := 0; i < numRows && k < n; i++ {
			res[i][x] = s[k]
			k++

		}
		x++

		for i := numRows - 2; i > 0 && k < n; i-- {

			res[i][x] = s[k]
			k++
			x++
		}

	}
	ans := ""
	for i := 0; i < numRows; i++ {
		ans += strings.Replace(string(res[i]), "\x00", "", -1)
	}
	return strings.Replace(ans, " ", "", -1)
}

func convert2(s string, numRows int) string {
	// 按行访问
	if numRows == 1 {
		return s
	}
	n := len(s)
	res := ""
	cyclelen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cyclelen {
			res += string(s[j+i])
			if i != 0 && i != numRows-1 && j+cyclelen-i < n {
				res += string(s[j+cyclelen-i])

			}
		}
	}
	return res
}

func convert3(s string, numRows int) string {
	// 按行访问
	n := len(s)
	if numRows == 1|| n<=numRows{
		return s
	}
	res :=[]byte{}
	grouplen := 2*numRows - 2
	var left,right int
	for i:=0;i<numRows;i++{
		left=i
		right=grouplen-left
		for {
			if left>=n{
				break
			}
			res=append(res,s[left])
			left+=grouplen
			if i!=0&&i!=numRows-1{
				if right>=n{
					break
				}
				res=append(res,s[right])

				right+=grouplen
			}
		}
		
	}
	
	return string(res)
}


func main() {
	fmt.Println(convert3("PAYPALISHIRING", 3), "LCIRETOESIIGEDHN")

}
