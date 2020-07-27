package main



func countDigitOne(n int) int {
	count:=0
	for i:=1;i<=n;i*=10{
		divider:=i*10
		count+=(n/divider)*i+min(max(n%divider-i+1,0),i)
	}
	return count

}

func max(x,y int) int {
	if x>y{
		return x
	}
	return y
}

func min(x,y int) int {
	if x<y{
		return x
	}
	return y
}

func main() {
	
}
