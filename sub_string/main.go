package main

import (
	"fmt"
)

// 无重复字符的最长子串
// 思路：

func LengthOfLongestSubstring(s string) int {
	var left, res int
	usedChar := make(map[rune]int)
	for i, v := range s {
		if k, ok := usedChar[v]; ok && left <= k {
			left = k + 1
		} else {
			res = max(res, i-left+1)
		}

		usedChar[v] = i
	}
	return res
}

func main() {
	s := "3456789x1234123"
	m := "sdkjfisf56789x1234123"
	fmt.Println(LengthOfLongestSubstring(s))
	fmt.Println(LengthOfLongestSubstring(m))
}
