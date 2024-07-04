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

	// 要当前字符的路
	process1(chs, i+1)
	tmp := chs[i]
	chs[i] = 0 // 将当前字符暂时设为0

	// 不要当前字符的路
	process1(chs, i+1)
	chs[i] = tmp // 恢复当前字符，以便在其他递归路径中正确使用
}

func printAllSubsequence2(str string) {
	chs := []rune(str)
	process2(chs, 0, make([]rune, 0, len(chs)))
}

// 当前来到i位置,要和不要，走两条路
// res之前的选择，所形成的列表
func process2(chs []rune, i int, res []rune) {
	if i == len(chs) {
		fmt.Println(string(res))
		return
	}
	// 选择要当前字符的路径
	res = append(res, chs[i])
	process2(chs, i+1, res)
	res = res[:len(res)-1] // 回溯

	// 选择不要当前字符的路径
	process2(chs, i+1, res) // 不要当前字符的路
}

func main() {
	test := "abc"
	printAllSubsequence1(test)
	printAllSubsequence2(test)
}
