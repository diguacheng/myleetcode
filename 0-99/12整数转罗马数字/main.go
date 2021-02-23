package main

import "fmt"

func intToRoman(num int) string {
	type ordermap struct{
		value int
		str string
	}
	res:=""
	om:=[]ordermap{{1000,"M"},{900,"CM"},{500,"D"},{400,"CD"},{100,"C"},{90,"XC"},{50,"L"},{40,"XL"},{10,"X"},{9,"IX"},{5,"V"},{ 4,"IV"},{1,"I"}}
	var count int
	for _,item:=range om{
		if num==0{
			break
		}
		
		count,num=num/item.value,num%item.value
		for count>0{
			res+=item.str
			count--
		}
	}
	return res
	

	

}


func main(){
	fmt.Println(intToRoman(1994))

}