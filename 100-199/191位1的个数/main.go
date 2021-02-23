package main 


func hammingWeight(num uint32) int {
	count:=0 
	for i := 0; i <32;i++{
		if num&1==1{
			count++
		}
		num=num>>1
	}
	return count
}

func hammingWeight1(num uint32) int {
	count:=0 
	for i := 0; i <32;i++{
		num=num&(num-1)
		if num!=0{
			count++
		}else{
			break
		}

	}
	return count
}


func main() {
	

}