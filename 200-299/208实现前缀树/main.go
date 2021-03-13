// package main

// type Trie struct {
// 	IsEnd bool
// 	Next  [26]*Trie
// }

// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	return Trie{}
// }

// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string) {
// 	for _, v := range word {
// 		if this.Next[v-'a'] == nil {
// 			this.Next[v-'a'] = &Trie{}
// 		}
// 		this = this.Next[v-'a']
// 	}
// 	this.IsEnd = true
// }

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	for _, v := range word {
// 		if this.Next[v-'a'] == nil {
// 			return false
// 		} else {
// 			this = this.Next[v-'a']
// 		}

// 	}
// 	return this.IsEnd == true
// }

// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	for _, v := range prefix {
// 		if this.Next[v-'a'] == nil {
// 			return false
// 		} else {
// 			this = this.Next[v-'a']
// 		}
// 	}
// 	return true
// }

// /**
//  * Your Trie object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Insert(word);
//  * param_2 := obj.Search(word);
//  * param_3 := obj.StartsWith(prefix);
//  */

package main

import "fmt"

type Trie struct {
	IsEnd bool
	Next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	n := len(word)
	if n == 0 {
		this.IsEnd = true
		return
	}
	if this.Next[word[0]-'a'] == nil {
		nt := Constructor()
		this.Next[word[0]-'a'] = &nt
	}
	if n >= 1 {
		this.Next[word[0]-'a'].Insert(word[1:])
	}

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.IsEnd
	}
	for i := 0; i < 26; i++ {
		if this.Next[i] != nil {
			if int(word[0]-'a') == i {
				if this.Next[i].Search(word[1:]) == true {
					return true
				}
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	for i := 0; i < 26; i++ {
		if this.Next[i] != nil && int(prefix[0]-'a') == i {
			return this.Next[i].StartsWith(prefix[1:])
		}
	}
	return false
}

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartsWith("app"))
	t.Insert("app")
	fmt.Println(t.Search("app"))

}
