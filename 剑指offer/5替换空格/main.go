package main

import "strings"

func replaceSpace(s string) string {
	if s == "" {
		return ""
	}
	res := []rune{}
	for _, v := range s {
		if v == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, v)
		}
	}
	return string(res)

}

func replaceSpace1(s string) string {
	if s == "" {
		return ""
	}
	temp := strings.Split(s, " ")
	return strings.Join(temp, "%20")

}

func replaceSpace2(s string) string {
	if s == "" {
		return ""
	}
	n := len(s)
	res := []byte{}
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			res = append(res, '%', '2', '0')

		} else {
			res = append(res, s[i])
		}
	}
	return string(res)

}

func main() {

}


/**
 * 
 * 请实现有重复数字的有序数组的二分查找。
 * 输出在数组中第一个大于等于查找值的位置，
 * 如果数组中不存在这样的数，则输出数组长度加一
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
*/

func upper_bound_( n int ,  v int ,  a []int ) int {

    left,right:=0,n
    var mid int
    for left<right{
        mid=left+(right-left)/2 
        if a[mid]>=v{
           right=mid
        }else{
           left=mid+1
        }
    }
    if left!=n&&a[left]>=v{
        return left+1
    }
    return n+1
}
