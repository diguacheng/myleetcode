package main

import "fmt"

func equationsPossible(equations []string) bool {
	l := make([]int, 26)
	for i := 0; i < 26; i++ {
		l[i] = i
	}
	for _, equation := range equations {
		if equation[1] == '=' {
			index1 := int(equation[0] - 'a')
			index2 := int(equation[3] - 'a')
			union(l, index1, index2)

		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			index1 := int(equation[0] - 'a')
			index2 := int(equation[3] - 'a')
			if find(l, index1) == find(l, index2) {
				return false
			}

		}
	}
	return true

}

func union(l []int, index1, index2 int) {
	l[find(l, index1)] = find(l, index2)

}
func find(l []int, index int) int {
	for l[index] != index {
		l[index] = l[l[index]]
		index = l[index]
	}
	return index
}


func equationsPossible1(equations []string) bool {
	l := make([]int, 26)
	for i := 0; i < 26; i++ {
		l[i] = i
	}
	var find func(i int)int
	find=func (i int)int{
		if l[i]!=i{
			return find(l[i])
		}
		return l[i]

	}
	union:=func(i,j int){
		l[find(i)]=find(j)
	}

	for _, equation := range equations {
		if equation[1] == '=' {
			index1 := int(equation[0] - 'a')
			index2 := int(equation[3] - 'a')
			if index1>index2{
				index1,index2=index2,index1
			}
			union(index1, index2)
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			index1 := int(equation[0] - 'a')
			index2 := int(equation[3] - 'a')
			if find( index1) == find(index2) {
				return false
			}
		}
	}
	return true
}

func main() {
	s:=[]string{"f==a","a==b","f!=e","a==c","b==e","c==f"}
	fmt.Println(equationsPossible(s))
}
