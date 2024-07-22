package main

import "fmt"

/*
岛问题
一个矩阵中只有0和1两种值，每个位置都可以和自己的上、下、左、右四个位置相连，如
果有一片1连在一起，这个部分叫做一个岛，求一个矩阵中有多少个岛？
【举例】
001010
111010
100100
000000
这个矩阵中有三个岛

【进阶】
如何设计一个并行算法解决这个问题
*/

func countIslands(m [][]int) int {
	if len(m) == 0 {
		return 0
	}
	N := len(m)
	M := len(m[0])
	res := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if m[i][j] == 1 {
				res++
				infect(m, i, j, N, M)
			}
		}
	}
	return res
}

func infect(m [][]int, i, j, N, M int) {
	if i < 0 || j < 0 || i >= N || j >= M || m[i][j] != 1 {
		return
	}
	//	i,j没越界 并且当前位置值是1
	m[i][j] = 2
	// 上下左右进行感染
	infect(m, i+1, j, N, M)
	infect(m, i, i-1, N, M)
	infect(m, i, j+1, N, M)
	infect(m, i, j-1, N, M)
}

func main() {
	m1 := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0},
	}
	fmt.Println(countIslands(m1))
}
