package main


func fractionToDecimal(numerator int, denominator int) string {
	var gcd func(a,b int)int
	gcd=func(a,b int)int{
		var temp int
		for b!=0{
			temp=b
			b=a%b
			a=temp
		}
		return a
	}
	var a,b int
	if numerator>=denominator{
		a=numerator/denominator
		numerator=numerator%denominator
	}else{
		a=0
	}
	m:=gcd(denominator,numerator)
	numerator=numerator/m
	denominator=denominator/m 
	list:=

}

func main(){

}