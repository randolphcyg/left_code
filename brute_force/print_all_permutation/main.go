package main

import (
	"fmt"
)

/*
打印一个字符串的全部排列，要求不要出现重复的排列
*/

func Permutation(str string) []string {
	var res []string
	if str == "" {
		return res
	}
	chs := []rune(str)
	process(chs, 0, &res)
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

	// 分支限界 去重 也可以在最外部洗数据 但会慢一些
	visited := make([]bool, 26) // 记录当前层已使用的字符 visited[0..25]
	for j := i; j < len(chs); j++ {
		if !visited[chs[j]-'a'] { // 检查字符是否已被使用
			visited[chs[j]-'a'] = true
			chs[i], chs[j] = chs[j], chs[i] // 交换
			process(chs, i+1, res)
			chs[i], chs[j] = chs[j], chs[i]
		}
	}
}

func main() {
	permutations := Permutation("abca")
	// [abca abac acba acab aacb aabc baca baac bcaa cbaa caba caab]
	fmt.Println(permutations)
}
