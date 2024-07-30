package main

import (
	"fmt"
	"math"
)

/*
题目描述
快递公司每日早晨，给每位快递员推送需要淡到客户手中的快递以及路线信息，快递员自己又查找了一些客户与客户之间的路线距离信息，请你依据这些信息，给快递员设计一条最短路径，告诉他最短路径的距离。

不限制快递包裹送到客户手中的顺序，但必须保证都送到客户手中；

用例保证一定存在投递站到每位客户之间的路线，但不保证客户与客户之间有路线，客户位置及投递站均允许多次经过；

所有快递送完后，快递员需回到投递站；

输入描述
首行输入两个正整数n、m.

接下来n行，输入快递公司发布的客户快递信息，格式为：客户id 投递站到客户之间的距离distance

再接下来的m行，是快递员自行查找的客户与客户之间的距离信息，格式为：客户1id 客户2id distance

在每行数据中，数据与数据之间均以单个空格分割规格:

0 <=n <= 10 0 <= m <= 10 0 < 客户id <= 1000 0 < distance <= 10000

输出描述
最短路径距离，如无法找到，请输出-1

示例1
输入：
2 1
1 1000
2 1200
1 2 300
输出：
2500

说明：
快递员先把快递送到客户1手中，接下来直接走客户1到客户2之间的直通线路，最后走投递站和客户2之间的路，回到投递站，距离为1000+300+1200= 2500

示例2
输入：
5 1
5 1000
9 1200
17 300
132 700
500 2300
5 9 400
输出：
9200
*/

const INF = math.MaxInt32

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// 初始化距离矩阵
	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, n+1)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}

	// 客户id和索引下标的对照表
	idxMap := make(map[int]int)

	// 初始化客户id到投递站(0)之间的距离
	for idx := 1; idx <= n; idx++ {
		var cid, distance int
		fmt.Scan(&cid, &distance)
		dist[0][idx] = distance
		dist[idx][0] = distance
		idxMap[cid] = idx
	}

	// 初始化客户与客户之间的距离
	for i := 0; i < m; i++ {
		var cid1, cid2, distance int
		fmt.Scan(&cid1, &cid2, &distance)
		idx1 := idxMap[cid1]
		idx2 := idxMap[cid2]
		dist[idx1][idx2] = distance
		dist[idx2][idx1] = distance
	}

	// Floyd-Warshall算法求出所有点之间的最短距离
	for k := 0; k <= n; k++ {
		dist[k][k] = 0 // 自己到自己的距离为0
		for i := 0; i <= n; i++ {
			for j := 0; j <= n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	// dp[state][last] 当前情况走过的最短距离
	// state表示已经投递的客户（指定二进制位为1表示已经投递），last表示上一次投递的客户
	dp := make([][]int, 1<<(n+1))
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	dp[1][0] = 0 // 初始状态，在投递站

	for state := 0; state < (1 << (n + 1)); state++ {
		for i := 0; i <= n; i++ {
			if (state>>i)&1 == 1 && dp[state][i] != INF { // 如果i已经投递且可达
				for last := 0; last <= n; last++ {
					if (state>>last)&1 == 0 { // 如果last还未投递
						newState := state | (1 << last)
						dp[newState][last] = min(dp[newState][last], dp[state][i]+dist[i][last])
					}
				}
			}
		}
	}

	minDist := INF
	for i := 1; i <= n; i++ {
		minDist = min(minDist, dp[(1<<(n+1))-1][i]+dist[i][0])
	}

	if minDist == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(minDist)
	}
}
