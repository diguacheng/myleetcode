package main

import "fmt"

func addBinary(a string, b string) string {
	la,lb,l:=len(a)-1,len(b)-1,len(b)-1
	if la>lb{
		l=la
	}
	res:=make([]byte,l+1)
	carry:=0
	for la>=0&&lb>=0{
		temp:=int(a[la]-'0')+int(b[lb]-'0')+carry
		switch temp {
		case 3:
			carry=1
			res[l]='1'
		case 2:
			carry=1
			res[l]='0'
		case 1:
			carry=0
			res[l]='1'
		case 0:
			carry=0
			res[l]='0'
		}
		la--
		lb--
		l--	
	}
	for la>=0{
		temp:=int(a[la]-'0')+carry
		switch temp{
		case 2:
			carry=1
			res[l]='1'
		case 1:
			carry=0
			res[l]='0'
		case 0:
			carry=0
			res[l]='0'
			break
		}
		la--
		l--
	}
	for lb>=0{
		temp:=int(b[lb]-'0')+carry
		switch temp{
		case 2:
			carry=1
			res[l]='1'
		case 1:
			carry=0
			res[l]='0'
		case 0:
			carry=0
			res[l]='0'
			break
		}
		lb--
		l--
	}
	if carry==1{
		res[l]='1'
	}
	for i:=0;i<len(res);i++{
		if res[i]=='1'{
			res=res[i:]
			break
		}
	}
	return string(res)





}



func main(){
	fmt.Println(addBinary("11","1"))

}