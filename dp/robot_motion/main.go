package main

import "fmt"

/*
机器人运动问题

排成一排的N个位置：1~N，N大于等于2。
开始时机器人在其中的M位置上，1<=M<=N。
机器人可以往左往右走，若来到1位置，下一步只能走到2；若来到N位置，下一步只能走到N-1。
规定机器人必须走K步，最终来到位置P，1<=P<=N，这样的方法有多少种。
入参：N, M, K, P，返回方法数目。

输入：
5 2 3 3
输出：
3
*/

// 递归方式
func ways1(N, M, K, P int) int {
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}

	// 总共N个位置，从M点出发，还剩K步，返回最终能到达P的方法数
	return walk(N, M, K, P)
}

/*
N : 位置为1 ~ N，固定参数
cur : 当前在cur位置，可变参数
rest : 还剩res步没有走，可变参数
P : 最终目标位置是P，固定参数
该函数的含义：只能在1~N这些位置上移动，当前在cur位置，走完rest步之后，停在P位置的方法数作为返回值返回
*/
func walk(N, cur, rest, P int) int {
	// 如果没有剩余步数了，当前的cur位置就是最后的位置
	// 如果最后的位置停在P上，那么之前做的移动是有效的
	// 如果最后的位置没在P上，那么之前做的移动是无效的
	if rest == 0 {
		if cur == P {
			return 1
		} else {
			return 0
		}
	}

	// 如果还有rest步要走，而当前的cur位置在1位置上，那么当前这步只能从1走向2
	// 后续的过程就是，来到2位置上，还剩rest-1步要走
	if cur == 1 { // 1 -> 2
		return walk(N, 2, rest-1, P)
	}
	// 如果还有rest步要走，而当前的cur位置在N位置上，那么当前这步只能从N走向N-1
	// 后续的过程就是，来到N-1位置上，还剩rest-1步要走
	if cur == N {
		return walk(N, N-1, rest-1, P)
	}

	// 如果还有rest步要走，而当前的cur位置在中间位置上，那么当前这步可以走向左，也可以走向右
	// 走向左之后，后续的过程就是，来到cur-1位置上，还剩rest-1步要走
	// 走向右之后，后续的过程就是，来到cur+1位置上，还剩rest-1步要走
	// 走向左、走向右是截然不同的方法，所以总方法数要都算上
	return walk(N, cur-1, rest-1, P) + walk(N, cur+1, rest-1, P)
}

// 动态规划: 使用二维数组存储中间结果，避免重复计算 空O(K * N)
func ways2(N, M, K, P int) int {
	// 参数无效直接返回0
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}

	// dp[rest][cur] 表示在剩余 rest 步数时，停留在位置 cur 的方法数
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	dp[0][P] = 1

	for rest := 1; rest <= K; rest++ {
		for cur := 1; cur <= N; cur++ {
			if cur == 1 {
				dp[rest][cur] = dp[rest-1][2]
			} else if cur == N {
				dp[rest][cur] = dp[rest-1][N-1]
			} else {
				dp[rest][cur] = dp[rest-1][cur-1] + dp[rest-1][cur+1]
			}
		}
	}

	return dp[K][M]
}

// 优化的动态规划: 使用一维数组存储中间结果，进一步优化空间复杂度 空O(N)
func ways3(N, M, K, P int) int {
	// 参数无效直接返回0
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}

	// dp[cur] 表示在当前步数下，停留在位置 cur 的方法数
	dp := make([]int, N+1)
	dp[P] = 1

	for i := 1; i <= K; i++ {
		leftUp := dp[1] // 左上角的值
		for j := 1; j <= N; j++ {
			tmp := dp[j]
			if j == 1 {
				dp[j] = dp[j+1]
			} else if j == N {
				dp[j] = leftUp
			} else {
				dp[j] = leftUp + dp[j+1]
			}
			leftUp = tmp
		}
	}

	return dp[M]
}

func main() {
	fmt.Println(ways1(5, 2, 3, 3))
	fmt.Println(ways2(5, 2, 3, 3))
	fmt.Println(ways3(5, 2, 3, 3))
}
