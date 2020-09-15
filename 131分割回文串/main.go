package main

import "fmt"

// 失败了 

func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return [][]string{}
	}
	if n == 1 {
		return [][]string{{s}}
	}
	check := func(s string) bool {
		n := len(s)
		if n == 1 {
			return true
		}
		l := n / 2
		for i := 0; i < l; i++ {
			if s[i] != s[n-1-i] {
				return false
			}
		}
		return true
	}
	res := make([][]string, 0)
	temp := make([]string, 0)

	var help func(already, temp []string)
	help = func(already, temp []string) {
		newTemp := []string{}
		newAlready := []string{}
		for i := 0; i < len(temp); i++ {
			str := temp[i]
			lenStr := len(str)
			if lenStr < 2 {
				continue
			}
			for j := 1; j < lenStr; j++ {
				newTemp = []string{}
				newAlready = []string{}
				if check(str[:j]) && check(str[j:]) {

					// 将需要切分的分为3块 其中
					newTemp = append(newTemp, temp[:i]...)
					newTemp = append(newTemp, []string{str[:j], str[j:]}...)
					newTemp = append(newTemp, temp[i+1:]...)

					res = append(res, append(already, newTemp...))
					newAlready = append(already, temp[:i]...)
					newAlready = append(newAlready, []string{str[:j], str[j:]}...)
					help(newAlready, temp[i+1:])
					newTemp = []string{}
					newAlready = []string{}
				}
			}

		}
	}
	flag := false
	if check(s) == true {
		res = append(res, []string{s})
		help([]string{}, []string{s})
		return res
	}
	for i := 1; i < n; i++ {
		temp = []string{}
		if check(s[:i]) && check(s[i:]) {
			flag = true
			temp = append(temp, s[:i])
			temp = append(temp, s[i:])
			res = append(res, temp)
			help([]string{}, temp)
			temp = []string{}
		}
	}
	if flag == false && n%2 == 1 {
		if check(s[:n/2]) && check(s[n/2+1:]) {
			temp = []string{s[:n/2], string(s[n/2]), s[n/2+1:]}
			res = append(res, temp)
			help([]string{}, temp)
			temp = []string{}
		}

	}
	return res
}

func partition1(s string) [][]string {
	n:=len(s)
	if n<=1{
		return [][]string{{s}}
	}
	res:=[][]string{}
	check := func(s string) bool {
		n := len(s)
		if n == 1 {
			return true
		}
		l := n / 2
		for i := 0; i < l; i++ {
			if s[i] != s[n-1-i] {
				return false
			}
		}
		return true
	}
	var help func(temp []string,i int)
	help = func(temp []string,i int){
		if i==n{
			ans:= make([]string,len(temp))
			copy(ans,temp)
			res = append(res, ans)
		}
		for j:=i;j<n;j++{
			if check(s[i:j+1]){
				temp=append(temp, s[i:j+1])
				help(temp,j+1)
				temp=temp[:len(temp)-1]
			}
		}
	}
	help([]string{},0)
	return res

}



func main() {
	// 回溯法 从 后面回溯 。
	s := "fff"
	fmt.Println(partition1(s))
}
