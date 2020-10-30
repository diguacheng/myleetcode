package main

import (
	"fmt"
	"strings")


func isValidSerialization(preorder string) bool {
	p:=strings.Split(preorder,",")
	n:=len(p)
	c:=1 
	for i:=0; i <n;i++{
		c-- 
		if c<0{
			return false
		}
		if p[i]!="#"{
			c+=2
		}
	}
	return c==0
}


func isValidSerialization1(preorder string) bool {
	//p:=strings.Split(preorder,",")
	n:=len(preorder)
	c:=1 
	var prev byte
	for i:=0; i <n;i++{
		if preorder[i]==','{
			c--
			if c<0{
				return false
			}
			if prev!='#'{
				c+=2
			}
		}
		prev=preorder[i]

	}
	if preorder[n-1]!='#'{
		c++ 
	}else{
		c--
	}
	return c==0
}


func main() {
	fmt.Println(isValidSerialization("9,#,92,#,#"))
	fmt.Println(isValidSerialization1("9,#,92,#,#"))

}