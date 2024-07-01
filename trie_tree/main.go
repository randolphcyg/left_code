package main

import "fmt"

// TrieNode 前缀树节点
type TrieNode struct {
	Pass  int
	End   int
	Nexts []*TrieNode
}

func NewTrieNode() *TrieNode {
	// Nexts[0] == null 没有走向'a'的路
	// Nexts[0] != null 有走向'a'的路
	// Nexts[25] != null 有走向'z'的路
	return &TrieNode{Pass: 0, End: 0, Nexts: make([]*TrieNode, 26)}
}

// Insert 插入
func (node *TrieNode) Insert(word string) {
	if word == "" {
		return
	}
	chs := []rune(word)
	node.Pass++
	index := 0
	for i := 0; i < len(word); i++ { // 从左往右遍历字符
		index = int(chs[i] - 'a') // 由字符，对应成走向哪条路
		if node.Nexts[index] == nil {
			node.Nexts[index] = NewTrieNode()
		}
		node = node.Nexts[index]
		node.Pass++
	}
	node.End++
}

// Search 单词word之前加入过几次
func (node *TrieNode) Search(word string) int {
	if word == "" {
		return 0
	}
	chs := []rune(word)
	index := 0
	for i := 0; i < len(chs); i++ {
		index = int(chs[i] - 'a')
		if node.Nexts[index] == nil {
			return 0
		}
		node = node.Nexts[index]
	}
	return node.End
}

// prefixNumber 有多少是以pre为前缀的
func (node *TrieNode) prefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	chs := []rune(pre)
	index := 0
	for i := 0; i < len(chs); i++ {
		index = int(chs[i] - 'a')
		if node.Nexts[index] == nil {
			return 0
		}
		node = node.Nexts[index]
	}
	return node.Pass

}

// delete 删除
func (node *TrieNode) delete(word string) {
	if node.Search(word) != 0 { // 确定树中确实加入过word 才删除
		chs := []rune(word)
		node.Pass--
		index := 0
		for i := 0; i < len(chs); i++ {
			index = int(chs[i] - 'a')
			node.Nexts[index].Pass--
			if node.Nexts[index].Pass == 0 {
				node.Nexts[index] = nil
				return
			}
			node = node.Nexts[index]
		}
		node.End--
	}
}

func main() {
	root := NewTrieNode()

	root.Insert("hello")
	root.Insert("hell")
	root.Insert("heaven")
	root.Insert("goodbye")

	fmt.Println("Search 'hello':", root.Search("hello"))       // Output: 1
	fmt.Println("Search 'hell':", root.Search("hell"))         // Output: 1
	fmt.Println("Search 'heaven':", root.Search("heaven"))     // Output: 1
	fmt.Println("Search 'goodbye':", root.Search("goodbye"))   // Output: 1
	fmt.Println("Search 'he':", root.Search("he"))             // Output: 0
	fmt.Println("Prefix 'he':", root.prefixNumber("he"))       // Output: 3
	fmt.Println("Prefix 'good':", root.prefixNumber("good"))   // Output: 1
	fmt.Println("Prefix 'hello':", root.prefixNumber("hello")) // Output: 1

	root.delete("hello")
	fmt.Println("After deleting 'hello'")
	fmt.Println("Search 'hello':", root.Search("hello"))     // Output: 0
	fmt.Println("Search 'hell':", root.Search("hell"))       // Output: 1
	fmt.Println("Prefix 'he':", root.prefixNumber("he"))     // Output: 2
	fmt.Println("Prefix 'good':", root.prefixNumber("good")) // Output: 1
}
