package main

import (
	"fmt"
)

/*
给定两个长度都为N的数组weights和values，weights[i]和values[i]分别代表 i号物品的重量和价值。
给定一个正数bag，表示一个载重bag的袋子，你装的物 品不能超过这个重量。返回你能装下最多的价值是多少?
*/

// maxValue1 通过递归计算背包问题的最大价值
func maxValue1(weights []int, values []int, bag int) int {
	return process1(weights, values, 0, 0, bag)
}

func process1(weights []int, values []int, i int, alreadyweight int, bag int) int {
	if alreadyweight > bag {
		return 0
	}
	if i == len(weights) {
		return 0
	}
	return max(
		process1(weights, values, i+1, alreadyweight, bag),
		values[i]+process1(weights, values, i+1, alreadyweight+weights[i], bag),
	)
}

// maxValue2 通过动态规划计算背包问题的最大价值
func maxValue2(c []int, p []int, bag int) int {
	dp := make([][]int, len(c)+1)
	for i := range dp {
		dp[i] = make([]int, bag+1)
	}

	for i := len(c) - 1; i >= 0; i-- {
		for j := bag; j >= 0; j-- {
			dp[i][j] = dp[i+1][j]
			if j+c[i] <= bag {
				dp[i][j] = max(dp[i][j], p[i]+dp[i+1][j+c[i]])
			}
		}
	}
	return dp[0][0]
}

func main() {
	weights := []int{3, 2, 4, 7}
	values := []int{5, 6, 3, 19}
	bag := 11
	fmt.Println(maxValue1(weights, values, bag))
	fmt.Println(maxValue2(weights, values, bag))
}
