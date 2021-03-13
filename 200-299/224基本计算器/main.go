package main

import (
	"fmt"	
	"strconv"
)


func calculate(s string)int {
	n:=len(s)
	st:=make([]byte,0)
	res:=make([]string,0)
	i:=0
	for i<n{
		if s[i]==' '{
			i++
			continue
		}
		if s[i]=='('||s[i]=='+'||s[i]=='-'{
			st=append(st, s[i])
			i++
		}else if s[i]==')'{
			for st[len(st)-1]!='('{
				res=append(res, string(st[len(st)-1]))
				st=st[:len(st)-1]
			}
			st=st[:len(st)-1]
			i++
		}else {
			numstr:=make([]byte,0)
			for i<n&&s[i]!='('&&s[i]!=')'&&s[i]!='+'&&s[i]!='-'{
				if s[i]==' '{
					i++
					continue
				}
				numstr=append(numstr, s[i])
				i++
			}
			res=append(res, string(numstr))
		}
	}
	for len(st)>0{
		res=append(res, string(st[len(st)-1]))
		st=st[:len(st)-1]
	}
	fmt.Println(res)
	v:=make([]int,0)
	for i:=0;i<len(res);i++{
		if res[i]=="+"||res[i]=="-"{
			num1:=v[len(v)-1]
			num2:=v[len(v)-2]
			v=v[:len(v)-2]
			if res[i]=="+"{
				v=append(v,num1+num2)
			}else{
				v=append(v,num2-num1)
			}
		}else{
			num,_:=strconv.Atoi(res[i])
			v=append(v,num)
		}
	}
	return v[0]
}



func calculate1(s string)int {
	ans:=0
    ops := []int{1}
    sign := 1
    n := len(s)
    for i := 0; i < n; {
        switch s[i] {
        case ' ':
            i++
        case '+':
            sign = ops[len(ops)-1]
            i++
        case '-':
            sign = -ops[len(ops)-1]
            i++
        case '(':
            ops = append(ops, sign)
            i++
        case ')':
            ops = ops[:len(ops)-1]
            i++
        default:
            num := 0
            for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
                num = num*10 + int(s[i]-'0')
            }
            ans += sign * num
        }
    }
    return ans
}


func main() {
	s := "2-1+2"
	fmt.Println( calculate(s))

}