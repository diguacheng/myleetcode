package main

func findTheDifference(s string, t string) byte {
    l:=[26]int{}
    n:=len(s)
    i:=0 
    for i<n{
        l[s[i]-'a']--
        l[t[i]-'a']++
        i++
    }
    l[t[i]-'a']++
    for i=0;i<26;i++{	
        if l[i]==1{
            return byte(i+'a')
        }
    }
    return ' '
}

func main(){

}