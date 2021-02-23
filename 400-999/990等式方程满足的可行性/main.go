package main

func equationsPossible(equations []string) bool {
	l:=make([]int,26)
	for i:=0;i<26;i++ {
		l[i]=i
	}
	for _,equation:=range equations{
		if equation[1]=='='{
			index1:=int(equation[0]-'a')
			index2:=int(equation[3]-'a')
			union(l,index1,index2)

		}
	}
	for _,equation:=range equations{
		if equation[1]=='!'{
			index1:=int(equation[0]-'a')
			index2:=int(equation[3]-'a')
			if find(l,index1)==find(l,index2){
				return false
			}

		}
	}
	return true

}

func union(l []int,index1,index2 int){
	l[find(l,index1)]=find(l,index2)

}
func find(l []int,index int)int{
	for l[index]!=index{
		l[index]=l[l[index]]
		index=l[index]
	}
	return index
}

func main(){

}