package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	len2 := len(s2)
	len1 := len(s1)
	if len2 < len1 {
		return false
	}
	need := [26]int{}
	window := [26]int{}
	for i := 0; i < len1; i++ {
		need[s1[i]-'a']++
	}
	check := func() bool {
		for i := 0; i < 26; i++ {
			if window[i] != need[i] {
				return false
			}
		}
		return true
	}
	match := 0 // 表示匹配的字符个数
	left, right := 0, 0
	var idx, leftidx byte
	for right < len2 {
		idx = s2[right] - 'a'
		window[idx]++
		right++
		if window[idx] <= need[idx] {
			match++
		}
		for match == len1 {
			if check() {
				return true
			}
			leftidx = s2[left] - 'a'
			left++
			if window[leftidx] <= need[leftidx] {
				match--
			}
			window[leftidx]--
		}
	}
	return false
}

func main() {
	s1 := "adc"
	s2 := "dcda"
	fmt.Println(checkInclusion(s1, s2))
}




// // 超时 
// func checkInclusion(s1 string, s2 string) bool {
//     m:=ToMap(s1)
//     n:=len(s1)
//     n2:=len(s2)
//     for i:=0;i<n2-n+1;i++{
//         if m[s2[i]]>0{
//             if check(m,s2[i:i+n]){
				
//                 return true
//             }
//         }
//     }
//     return false
// }

// func check(m map[byte]int,s string)bool{
//     m1:=ToMap(s)
//     for k,v:=range m{
//         if m1[k]!=v{
//             return false
//         }
//     }
//     return true
// }

// func ToMap(s string)map[byte]int{
//     m:=make(map[byte]int)
//     for i:=0;i<len(s);i++{
//         m[s[i]]++
//     }
//     return m
// }