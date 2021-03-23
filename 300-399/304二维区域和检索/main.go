package main
type NumMatrix struct {
    m [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m:=len(matrix)
    if m==0{
        return NumMatrix{nil}
    }
    n:=len(matrix[0])
    if n==0{
        return NumMatrix{nil}
    }
    arr:=make([][]int,m)
    for i:=0;i<m;i++{
        arr[i]=make([]int,n)
        copy(arr[i],matrix[i])
        for j:=1;j<n;j++{
            arr[i][j]+=arr[i][j-1]
        }
    }
    for i:=0;i<n;i++{
        for j:=1;j<m;j++{
            arr[j][i]+=arr[j-1][i]
        }
    }
    return NumMatrix{m:arr}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    if row1==0&&col1==0{
        return this.m[row2][col2]
    }
    if row1==0{
        return this.m[row2][col2]-this.m[row2][col1-1]
    }
    if col1==0{
        return this.m[row2][col2]-this.m[row1-1][col2]
    }
    return this.m[row2][col2]-this.m[row1-1][col2]-this.m[row2][col1-1]+this.m[row1-1][col1-1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */