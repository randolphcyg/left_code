package main

import "fmt"

/*
打印一个字符串的全部子序列，包括空字符串
*/

func printAllSubsequence1(str string) {
	chs := []rune(str)
	process1(chs, 0)
}

func process1(chs []rune, i int) {
	if i == len(chs) {
		fmt.Println(string(chs))
		return
	}
	process1(chs, i+1) // 要当前字符的路
	tmp := chs[i]
	chs[i] = 0
	process1(chs, i+1) // 不要当前字符的路
	chs[i] = tmp
}

func printAllSubsequence2(str string) {
	chs := []rune(str)
	process2(chs, 0, []rune{})
}

// 当前来到i位置,要和不要，走两条路
// res之前的选择，所形成的列表
func process2(chs []rune, i int, res []rune) {
	if i == len(chs) {
		printList(res)
		return
	}
	resKeep := copyList(res)
	resKeep = append(resKeep, chs[i])
	process2(chs, i+1, resKeep) // 要当前字符的路
	resNoInclude := copyList(res)
	process2(chs, i+1, resNoInclude) // 不要当前字符的路
}

func printList(res []rune) {
	for _, ch := range res {
		fmt.Print(string(ch))
	}
	fmt.Println()
}

func copyList(list []rune) []rune {
	newList := make([]rune, len(list))
	copy(newList, list)
	return newList
}

func main() {
	test := "abc"
	printAllSubsequence1(test)
	printAllSubsequence2(test)
}
