package main

import "fmt"

/*
规定1和A对应、2和B对应、3和C对应... 那么一个数字字符串比如"111"，就可以转化为"AAA"、"KA"和"AK"。
给定一个只有数字字符组成的字符串str，返回有多少种转化结果。
*/

func number(str string) int {
	if str == "" {
		return 0
	}
	return process([]rune(str), 0)
}

func process(chs []rune, i int) int {
	if i == len(chs) {
		return 1
	}
	if chs[i] == '0' {
		return 0
	}
	if chs[i] == '1' {
		res := process(chs, i+1) // i自己作为单独的部分，后续有多少种方法
		if i+1 < len(chs) {
			res += process(chs, i+2) // i和i+1作为单独的部分
		}
		return res
	}
	if chs[i] == '2' {
		res := process(chs, i+1)
		// i和i+1作为单独的部分并且没有超过26，后续有多少种方法
		if i+1 < len(chs) && (chs[i+1] >= '0' && chs[i+1] <= '6') {
			res += process(chs, i+2) // i和i+1作为单独的部分
		}
		return res
	}
	// chs[i] == '3' ~ '9'
	return process(chs, i+1)
}

func main() {
	fmt.Println(number("11111"))
}
