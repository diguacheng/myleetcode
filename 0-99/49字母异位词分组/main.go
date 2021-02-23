package main

import (
	"fmt"
	"sort"
)

// 无法解决有重复字母
func groupAnagrams(strs []string) [][]string {
	table := make(map[int][]string)
	for _, str := range strs {
		x := 0
		for i := 0; i < len(str); i++ {
			x = x | 1<<int(str[i]-'a')
		}

		table[x] = append(table[x], str)
	}

	res := make([][]string, len(table))
	i := 0
	for _, v := range table {
		res[i] = make([]string, len(v))
		copy(res[i], v)
		i++
	}
	return res
}

// 构建一个哈希表 key 为 string类型，为每个字符串作为[]byte排序后的值
// valu为 []string 储存 具有相同key的strig
func groupAnagrams1(strs []string) [][]string {
	table := make(map[string][]string)
	var temp []byte
	for _, str := range strs {
		temp = []byte(str)
		sort.Slice(temp, func(i, j int) bool {
			return byte(temp[i]) < byte(temp[j])
		})
		table[string(temp)] = append(table[string(temp)], str)
	}

	res := make([][]string, len(table))
	i := 0
	for _, v := range table {
		res[i] = make([]string, len(v))
		copy(res[i], v)
		i++
	}
	return res
}

func ssort() {
	s := []byte("wertgh")
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(string(s))
}

func main() {
	strs := []string{
		//"ddddddddddg","dgggggggggg",
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	fmt.Println(groupAnagrams1(strs))
	ssort()
}
