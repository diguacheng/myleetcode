package main

//TODO 未完成
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1==0{
		return 0
	}
	s1cnt,index,s2cnt:=0,0,0
	recall:=make(map[int]interface{})
	for {
		s1cnt++
		for i:=0;i<len(s1);i++{
			if s1[i]==s2[index]{
				index++
				if index==len(s2){
					s2cnt++
					index=0
				}
			}
		}
		if s1cnt==n1{
			return s2cnt/n2
		}
		if 
	}


}

func main() {

}