package main

import "fmt"

/*
分月饼

题目描述
中秋节，公司分月饼，m个员工，买了n个月饼，m <= n，每个员工至少分1个月饼，但可以分多个，
单人分到最多月饼的个数是Max1，单人分到第二多月饼个数是Max2，Max1-Max2 <= 3，
单人分到第n-1多月饼个数是Max(n-1)，单人分到第n多月饼个数是Max(n)，Max(n-1)- Max(n) <= 3，
问有多少种分月饼的方法?

输入描述
每一行输入m n，表示m个员工，n个月饼，m<=n

输出描述
输出有多少种月饼分法

示例
输入
2 4
输出
2
说明
分法有2种:

4=1+3
4=2+2

注意: 1+3和3+1算一种分法

输入
3 12
输出
6
*/

// 动态规划函数，用于计算分月饼的方法数
func distribute(employees, cakes int) int {
	remainingCakes := cakes - employees
	// dp是一个三维数组，其中 dp[i][j][k] 表示将 i 个剩余月饼分配给 j 个员工，且最多有一个员工得到 k 个月饼的分配方案数
	dp := make([][][]int, remainingCakes+1)
	for i := range dp {
		dp[i] = make([][]int, employees+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, remainingCakes+1)
		}
	}

	// 初始化 dp 数组
	for i := 0; i <= remainingCakes; i++ {
		dp[i][1][i] = 1
	}

	// 动态规划填表
	for i := 0; i <= remainingCakes; i++ { // 外层循环 i 遍历所有可能的剩余月饼数
		for j := 2; j <= employees; j++ { // 中间循环 j 遍历员工数，从2开始，因为1的情况已经初始化
			for k := 0; k <= i; k++ { // 内层循环 k 遍历分配给单个员工的月饼数
				dp[i][j][k] = dp[i-k][j-1][k] // 最多一个员工得到k个
				if k+1 <= i {
					dp[i][j][k] += dp[i-k-1][j-1][k+1] // 存在一个员工得到k+1个
				}
				if k+2 <= i {
					dp[i][j][k] += dp[i-k-2][j-1][k+2]
				}
				if k+3 <= i {
					dp[i][j][k] += dp[i-k-3][j-1][k+3]
				}
			}
		}
	}

	sum := 0
	for _, value := range dp[remainingCakes][employees] {
		sum += value
	}

	return sum
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	fmt.Println(distribute(m, n))
}
