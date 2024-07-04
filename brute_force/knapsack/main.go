package main

import (
	"fmt"
	"math"
)

/*
背包问题
给定两个长度都为N的数组weights和values，weights[i]和values[i]分别代表i号物品的重量和价值。
给定一个正数bag，表示一个载重bag的袋子，你装的物品不能超过这个重量。返回你能装下最多的价值是多少?
*/

// 暴力解法 递归
func maxValue1(weights []int, values []int, bag int) int {
	return process1(weights, values, 0, 0, bag)
}

func process1(weights []int, values []int, i int, alreadyWeight int, bag int) int {
	if alreadyWeight > bag {
		return math.MinInt32 // 使用负无穷表示无效选择
	}
	if i == len(weights) {
		return 0
	}
	return max(
		process1(weights, values, i+1, alreadyWeight, bag),
		values[i]+process1(weights, values, i+1, alreadyWeight+weights[i], bag),
	)
}

// 动态规划
func maxValue2(c []int, p []int, bag int) int {
	dp := make([][]int, len(c)+1)
	for i := range dp {
		dp[i] = make([]int, bag+1)
	}

	for i := len(c) - 1; i >= 0; i-- {
		for j := 0; j <= bag; j++ {
			// 不选择当前物品
			dp[i][j] = dp[i+1][j]
			// 选择当前物品（前提是当前物品可以放入剩余容量中）
			if j >= c[i] {
				dp[i][j] = max(dp[i][j], p[i]+dp[i+1][j-c[i]])
			}
		}
	}
	return dp[0][bag]
}

func main() {
	weights := []int{3, 2, 4, 7}
	values := []int{5, 6, 3, 19}
	bag := 11
	/*
		25
		25
	*/
	fmt.Println(maxValue1(weights, values, bag))
	fmt.Println(maxValue2(weights, values, bag))
}
