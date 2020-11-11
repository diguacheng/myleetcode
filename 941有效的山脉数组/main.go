package main 


func validMountainArray(A []int) bool {
	n:=len(A)
	if n<3{
		return false
	}
	i:=1
	for i<n&&A[i-1]<A[i]{
		i++
	}
	if i==1||i==n{
		return false
	}
	for i<n&&A[i-1]>A[i]{
		i++
	}
	if i==n{
		return true
	}
	return false

}

func main(){

}