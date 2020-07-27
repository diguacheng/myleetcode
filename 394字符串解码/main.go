package main

import "strconv"

func decodeString(s string) string {

	res:=""

	nums:=make([]int,0)
	stack:=make([]byte,0)
	for i<=len(s) {
		if s[i]>='0'&&s[i]<='9'{
			j:=i
			for ;s[i+1]>='0'&&s[i+1]<='9';i++{}
			num,_:=strconv.Atoi(s[j:i])
			nums=append(nums,num)
			continue
		}
		if s[i]==']'{
			for j:=i;s[j]!='[';j--{
				
			}
			

		}
		if s[i]=='['{
			stack=append(stack,s[i])
		}
		

			

		}
	}
}



func main(){

}