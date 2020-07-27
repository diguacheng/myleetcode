package main

import "fmt"

func exist(board [][]byte, word string) bool {
	isVisited:=make([][]bool,len(board))
	for i:= range isVisited{
		isVisited[i]=make([]bool,len(board[0]))
	}
	for i:=range board{
		for j:=range board[0]{
			if haspath(board,i,j,word,0,isVisited){
				return true
			}

		}
	}
	return false
}

func haspath(board [][]byte,i int,j int,word string,length int,isVisited [][]bool)bool{
	if length==len(word){
		return true
	}
	flag:=false
	if i>=0&&i<len(board) && j>=0&&j<len(board[0])&&board[i][j]==word[length]&&!isVisited[i][j]{
		length++
		isVisited[i][j]=true
		flag=haspath(board,i-1,j,word,length,isVisited)||haspath(board,i,j-1,word,length,isVisited)||haspath(board,i+1,j,word,length,isVisited)||haspath(board,i,j+1,word,length,isVisited)
		if !flag{
			length--
			isVisited[i][j]=false
		}
	}
	return flag
}

func main() {
	board:= [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word:="ABCCED"
	fmt.Println(exist(board,word))

}