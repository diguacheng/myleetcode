package main

func findCircleNum(isConnected [][]int) int {
	m:=len(isConnected)
	isVisited:=make([][]int,m)
	for i:=0;i<m;i++{
		isVisited[i]=make([]int,m)
	}
	cnt:=0
	for i:=0;i<m;i++{
		for j:=0;j<m;j++{
			if isVisited[i][j]==0&&isConnected[i][j]==1{
				isVisited[i][j]=1
				isVisited[j][i]=1
				cnt++
				DFS(isVisited,isConnected,i,j,m)
			}
		}
	}
	return cnt
}

func DFS(isVisited,isConnected [][]int, i,j,m int){
	for y:=0;y<m;y++{
		if isVisited[j][y]==0&&isConnected[j][y]==1{
			isVisited[j][y]=1
			isVisited[y][j]=1
			DFS(isVisited,isConnected,j,y,m)
		}
	}
	for y:=j+1;y<m;y++{
		if isVisited[i][y]==0&&isConnected[i][y]==1{
			isVisited[i][y]=1
			isVisited[y][i]=1
			DFS(isVisited,isConnected,i,y,m)
		}
	}
}


func main(){

}