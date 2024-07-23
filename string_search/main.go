package main

/*
暴力子字符串查找
找到模式子串就返回true和索引，
不匹配就返回false和文本长度
*/
func bruteForceStringMatching(txt string, pat string) (bool, int) {
	m := len(pat)
	n := len(txt)

	for i := 0; i < n-m; i++ {
		j := 0
		for j = 0; j < m; j++ {
			if txt[i+j] != pat[j] {
				break
			}
		}
		if j == m {
			return true, i // 找到匹配
		}
	}

	return false, n //未找到匹配
}

// 暴力子字符串查找（显式回退）
func bruteForceStringMatching2(txt string, pat string) (bool, int) {
	m, j := len(pat), len(pat)
	n, i := len(txt), len(txt)

	for i, j = 0, 0; i < n && j < m; i++ {
		if txt[i] == pat[j] {
			j++
		} else {
			i -= j
			j = 0
		}
	}

	if j == m {
		return true, i - m // 找到匹配
	}
	return false, n //未找到匹配
}

func main() {
	println(bruteForceStringMatching("this is a test text", "hello"))     // false 19
	println(bruteForceStringMatching("this is a hello message", "hello")) // true 10

	println(bruteForceStringMatching2("this is a test text", "hello"))     // false 19
	println(bruteForceStringMatching2("this is a hello message", "hello")) // true 10
}
