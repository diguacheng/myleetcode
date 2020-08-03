package main

import (
	"fmt"
	"strconv")


func evalRPN(tokens []string) int {
	numStack := make([]int, 0)
	res:=0
	for _, v := range tokens {
		num,ok:=strconv.Atoi(v)
		if ok==nil{

			numStack = append(numStack, num)
		}else{
			num1:=numStack[len(numStack)-1]
			num2:=numStack[len(numStack)-2]
			numStack=numStack[:len(numStack)-2]
			switch v{
			case "+":
				res=num1+num2
			case "-":
				res=num2-num1
			case "*":
				res=num1*num2
			case "/":
				res=num2/num1
			}
			numStack = append(numStack, res)
		}
	}
	return numStack[0]

}

func main() {
	tokens:=[]string{"2","1","+","3","*"}
	fmt.Println(evalRPN(tokens))


}