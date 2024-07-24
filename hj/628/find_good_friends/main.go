package main

import (
	"fmt"
)

/*
题目描述
在学校中，N个小朋友站成一队， 第i个小朋友的身高为height[i]，第i个小朋友可以看到的右边的第一个比自己身高更高的小朋友j，
那么j是i的好朋友（j > i）。请重新生成一个列表，对应位置的输出是每个小朋友的好朋友位置，
如果没有看到好朋友，请在该位置用0代替。小朋友人数范围是 [0, 40000]。

输入描述
第一行输入N，表示有N个小朋友

第二行输入N个小朋友的身高height[i]，都是整数

输出描述
输出N个小朋友的好朋友的位置

示例一

输入
2
100 95
输出
0 0

示例二

输入
8
123 124 125 121 119 122 126 123
输出
1 2 6 5 5 6 0 0
*/

// findFriends 找到每个小朋友的好朋友位置
func findFriends(heights []int) (result []int) {
	size := len(heights)
	result = make([]int, size)

	for i := 0; i < size; i++ {
		if i == size-1 {
			result[i] = 0
			break
		}
		for j := i + 1; j < size; j++ {
			if heights[j] > heights[i] {
				result[i] = j
				break
			}
			if j == size-1 {
				result[i] = 0
				break
			}
		}
	}

	return
}

// findFriends 找到每个小朋友的好朋友位置
func findFriends2(heights []int) (result []int) {
	size := len(heights)
	result = make([]int, size)
	var stack []int // 储存小朋友的索引

	for i := 0; i < size; i++ {
		// 内层循环检查栈顶元素对应的小朋友的身高是否小于当前小朋友的身高：
		// 如果是，将当前小朋友的索引 i 作为栈顶小朋友的好朋友位置，并将栈顶小朋友出栈。
		// 继续检查下一个栈顶元素，直到栈为空或栈顶元素不再小于当前小朋友的身高。
		for len(stack) > 0 && heights[stack[len(stack)-1]] < heights[i] {
			result[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i) // 当前小朋友压入栈
	}

	// 栈中剩余的元素没有好朋友
	for len(stack) > 0 {
		result[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1]
	}

	return result
}

/*
输入
8
123 124 125 121 119 122 126 123
输出
1 2 6 5 5 6 0 0
*/
func main() {
	var N int
	fmt.Scan(&N)

	if N == 0 {
		return
	}

	heights := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&heights[i])
	}

	fmt.Println("双层循环方法：")
	result := findFriends(heights)

	for _, h := range result {
		fmt.Print(h, " ")
	}
	fmt.Println()

	fmt.Println("单调栈方法：")
	result2 := findFriends2(heights)

	for _, h := range result2 {
		fmt.Print(h, " ")
	}
	fmt.Println()
}
