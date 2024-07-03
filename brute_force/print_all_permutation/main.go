package main

import (
	"fmt"
	"slices"
)

/*
打印一个字符串的全部排列
打印一个字符串的全部排列，要求不要出现重复的排列
*/

func Permutation(str string) []string {
	var res []string
	if str == "" {
		return res
	}
	chs := []rune(str)
	process(chs, 0, &res)
	slices.Sort(res)
	return res
}

// str[i..]范围上，所有的字符，都可以在i位置上，后续都去尝试
// str[0..i-1]范围上，是之前做的选择
// 请把所有字符串形成的全排列，加到res中
func process(chs []rune, i int, res *[]string) {
	if i == len(chs) {
		*res = append(*res, string(chs))
		return
	}
	visit := make([]bool, 26)
	for j := i; j < len(chs); j++ {
		if !visit[chs[j]-'a'] {
			visit[chs[j]-'a'] = true
			chs[i], chs[j] = chs[j], chs[i]
			process(chs, i+1, res)
			chs[i], chs[j] = chs[j], chs[i]
		}
	}
}

func main() {
	test := "abc"
	permutations := Permutation(test)
	for _, perm := range permutations {
		fmt.Println(perm)
	}
}
