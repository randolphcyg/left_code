package main

import "fmt"

/*
KMP算法
字符串str1和str2，str1是否包含str2，如果包含返回str2在str1中开始的位置。 如何做到时间复杂度O(N)完成?
*/

func getIndexOf(str, patt string) int {
	if str == "" || patt == "" || len(str) < 1 || len(patt) < 1 {
		return -1
	}

	next := getNextArray(patt)

	i, j := 0, 0 // 主串中的指针、子串中的指针
	for i < len(str) && j < len(patt) {
		if str[i] == patt[j] { // 字符匹配 指针后移
			i++
			j++
		} else if j > 0 { // 字符失配 根据next跳过子串前面的一些字符
			j = next[j]
		} else { // 字符第一个字符就失配
			i++
		}
	}

	if j == len(patt) { // 匹配成功
		return i - j
	}
	return -1
}

// next数组
func getNextArray(patt string) []int {
	if len(patt) == 1 {
		return []int{-1}
	}

	next := make([]int, len(patt)) // next数组
	prefixLen := 0                 // 当前共同前后缀长度
	i := 1

	for i < len(patt) {
		if patt[prefixLen] == patt[i] {
			prefixLen++
			next[i] = prefixLen
			i++
		} else if prefixLen > 0 {
			prefixLen = next[prefixLen-1]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}

func main() {
	str := "abcabcbyd汉dmiccc"
	match := "byd汉dmi"
	fmt.Println(getIndexOf(str, match))
}
