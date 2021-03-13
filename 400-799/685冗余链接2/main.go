package main

func findRedundantDirectedConnection(edges [][]int) []int {
	n:=len(edges)
	p:=make([]int,n+1)
	for i:=1;i<n;i++{
		p[i]=i
	}
	isCircle:=func(i int)bool{
		if p[i]==i{
			return false
		}
		v:=i
		i=p[i]
		for p[i]!=v{
			if p[i]==i{
				return false
			}
			i=p[i]
		}
		return true
	}
	var find func(i int)int
	find=func(i int) int {
		if p[i]!=i{
			return find(p[i])
		}
		return p[i]
	}
	var union func(i,j int)
	union=func(i,j int){
		p[find(i)]=find(j)
	}
	for i:=0;i<n;i++{
		s,e:=edges[i][0],edges[i][1]
		union(s,e)
		if !isCircle(s){
			return edges[i]
		}
	}
	return nil
}


func main(){

}