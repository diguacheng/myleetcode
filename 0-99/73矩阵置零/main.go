package main 

func setZeroes(matrix [][]int)  {
	n,m:=len(matrix), len(matrix[0])
	var isCol bool
	for i:=0;i<n;i++{
		if matrix[i][0]==0{
			isCol = true
		}
		for j:=1;j<m;j++{
			if matrix[i][j]==0{
				matrix[0][j]=0
				matrix[i][0]=0
			}
		}
	}

	for i:=1;i<n;i++{
		for j:=1;j<m;j++{
			if matrix[i][0]==0||matrix[0][j]==0{
				matrix[i][j]=0
			}
		}
	}
	if matrix[0][0]==0{
		for j:=1;j<m;j++{
			matrix[0][j]=0
		}
	}
	if isCol{
		for i:=0;i<n;i++{
			matrix[i][0]=0
		}
	}
}

func main() {

}