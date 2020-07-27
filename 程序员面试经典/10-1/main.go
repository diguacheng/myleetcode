package main

func merge(A []int, m int, B []int, n int)  {
	i:=m+n
	for n>0 {
		if A[m-1]>B[n-1]{
			A[i-1]=A[m-1]
			i--
			m--
		}else{
			A[i-1]=B[n-1]
			i--
			n--
		}
	}
	for n>0{
		A[i-1]=B[n-1]
		i--
		n--
	}



}

func main(){

}