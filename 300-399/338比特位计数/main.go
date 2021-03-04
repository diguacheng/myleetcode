package main

import "fmt"

func countBits(num int) []int {
	if num == 0 {
		return []int{0}

	}
	res := make([]int, num+1)
	res[1] = 1
	a := 1
	index := 2
	for index <= num {
		a = a << 1
		for i := 0; i < a && index <= num; i++ {
			res[index] = res[i] + 1
			index++
		}

	}
	return res

}


func countBits1(num int) []int {
    if num==0{
        return []int{0}
    }
    res:=[]int{0,1}
    i:=2
    for i<=num{
        cnt:=0
        curr:=i 
        for curr>0{
            cnt+=curr&1
            curr=curr>>1
        }
        res=append(res,cnt)
		i++
    }
    return res
}

func main() {
	fmt.Println(countBits1(1))
	fmt.Println(countBits1(5))

}
