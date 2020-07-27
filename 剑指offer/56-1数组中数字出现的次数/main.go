package main

import "fmt"

func singleNumbers(nums []int) []int {
	res := 0
	for _, num := range nums {
		// 相同的数异或值为0，不同的数异或值为1  
		res = res ^ num // 得到两个不同的数的异或值
	}
	div := 1
	for div&res == 0 {
		div = div << 1// 找到异或值某一位为1，作为分类的依据
	}
	a, b := 0, 0
	for _, num := range nums {
		//fmt.Println(div&num,div,num)
		//fmt.Printf("%b %b %b %d \n",div,num,div&num,div&num)
		if div&num!=0{
			//  相同的数异或为0，a^b=1,说明a、b某位是不同的，所以a、b的与值不同，可分为两组
			a = a ^ num
		} else {
			b = b ^ num
		}

	}
	return []int{a, b}

}

func main() {
	nums := []int{1,2,5,2}
	fmt.Println(singleNumbers(nums))

}
