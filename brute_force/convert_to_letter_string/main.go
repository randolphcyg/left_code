package main

import "fmt"

/*
规定'1'对应'A'、'2'对应'B'、'3'对应'C'... '26'对应'Z'
那么一个数字字符串比如"111"，就可以转化为"AAA"、"KA"和"AK"。
给定一个只有数字字符组成的字符串str，返回有多少种转化结果。
*/

// 暴力解法 递归
func number1(str string) int {
	if str == "" {
		return 0
	}
	return process([]rune(str), 0)
}

// 递归
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

/*
状态转移方程
如果当前字符是 '0'，则不能单独解码，dp[i] = 0。
如果当前字符是 '1'，则 dp[i] = dp[i+1]（单独解码为 'A'）+ dp[i+2]（和后一个字符一起解码）。
如果当前字符是 '2'，则需要进一步判断：
	如果后一个字符是 '0' 到 '6'，则 dp[i] = dp[i+1] + dp[i+2]。
	否则，只能单独解码，dp[i] = dp[i+1]。
对于其他字符（'3' 到 '9'），只能单独解码，dp[i] = dp[i+1]。
*/

// 推荐：动态规划 时O(n)空O(n)
func number2(str string) int {
	if str == "" {
		return 0
	}
	chs := []rune(str)
	n := len(chs)
	dp := make([]int, n+1)
	dp[n] = 1 // 空字符串有一种解码方式

	for i := n - 1; i >= 0; i-- {
		if chs[i] == '0' {
			dp[i] = 0
		} else if chs[i] == '1' {
			dp[i] = dp[i+1]
			if i+1 < n {
				dp[i] += dp[i+2]
			}
		} else if chs[i] == '2' {
			dp[i] = dp[i+1]
			if i+1 < n && chs[i+1] >= '0' && chs[i+1] <= '6' {
				dp[i] += dp[i+2]
			}
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}

// 新增辅助函数来打印所有可能的转换结果
func printAllConversions(str string) {
	if str == "" {
		return
	}
	chs := []rune(str)
	processWithPrint(chs, 0, []rune{})
}

func processWithPrint(chs []rune, i int, current []rune) {
	if i == len(chs) {
		fmt.Println(string(current))
		return
	}
	if chs[i] == '0' {
		return
	}
	// 处理单个字符
	current = append(current, rune('A'+(chs[i]-'1')))
	processWithPrint(chs, i+1, current)
	current = current[:len(current)-1]

	// 处理两个字符的组合
	if i+1 < len(chs) && (chs[i] == '1' || (chs[i] == '2' && chs[i+1] >= '0' && chs[i+1] <= '6')) {
		num := (chs[i]-'0')*10 + (chs[i+1] - '0')
		current = append(current, rune('A'+(num-1)))
		processWithPrint(chs, i+2, current)
		current = current[:len(current)-1]
	}
}

func main() {
	/*
		3
		3
		ABC
		AW
		LC
	*/
	s1 := "123"
	fmt.Println(number1(s1))
	fmt.Println(number2(s1))
	printAllConversions(s1)
}
