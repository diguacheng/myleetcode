package main

import "fmt"

func countCharacters(words []string, chars string) int {
	charsmap := make(map[rune]int)
	for _, char := range chars {
		charsmap[char]++
	}
	res := 0
	for _, word := range words {
		wordmap := make(map[rune]int)
		for _, char := range word {
			wordmap[char]++
		}
		isres := true
		for _, char := range word {
			if charsmap[char] < wordmap[char] {
				isres = false
				break
			}
		}
		if isres {
			res = res + len(word)
		}

	}
	return res

}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Println(countCharacters(words, chars))

}
