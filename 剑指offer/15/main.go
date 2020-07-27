package main

import "fmt"

func hammingWeight(n int) int {
	count :=0
	flag:=1
	for flag>0{
		if n&flag>0{

			count++
		}
		flag=flag<<1
	}
	return count

}

func hammingWeight1(n int) int {
	count := 0
	for n>0{
		count++
		n=(n-1)&n
	}
	return count

}


func main() {
	fmt.Println(hammingWeight1(00000000000000000000000000001011))

}