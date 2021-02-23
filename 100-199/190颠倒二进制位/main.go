package main

import "fmt"

func reverseBits(num uint32) uint32 {
	res:=0
	i:=0
	for num>0{
		if num%2==1{
			res+=1<<(31-i)
		}
		i++
		num=num>>1
	}
	return uint32(res)
}
func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(1<<31)

}