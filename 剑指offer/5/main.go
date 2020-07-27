package main



func replaceSpace(s string) string {
	if s==""{
		return ""
	}
	res:=[]rune{}
	for _,v:=range s{
		if v == ' '{
			res=append(res,'%','2','0')
		}else{
			res=append(res,v)
		}
	}
	return string(res) 
    
}


func main() {


}