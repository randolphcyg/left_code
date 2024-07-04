package main

import (
	"fmt"
	"math"
	"time"
)

/*
N皇后问题是指在N*N的棋盘上要摆N个皇后，要求任何两个皇后不同行、不同列，也不在同一条斜线上。

给定一个整数n，返回n皇后的摆法有多少种。
n=1，返回1。

n=2或3，2皇后和3皇后问题无论怎么摆都不行，返回0。n=8，返回92。
*/

/*
复杂度 O(n^n) 无法优化
常数优化
*/

// 暴力解法 递归
func nQueen1(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n) // record[i] -> i行的皇后放在了第几列
	return process1(0, record, n)
}

// i 当前第i行
// record[0..i-1]表示之前的行，放了的皇后位置
// n表示整体一共有多少行
// 返回值是，摆完所有的皇后，合理的摆法有多少种
func process1(i int, record []int, n int) int {
	if i == n { // 终止行
		return 1
	}
	res := 0
	for j := 0; j < n; j++ { // 尝试i行所有的列j
		// 当前i行的皇后，放在j列，会不会和之前的0..i-1的皇后，共行共列共斜线
		if isValid(record, i, j) {
			record[i] = j
			res += process1(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		// math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k)) 共斜线
		if j == record[k] || math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k)) {
			return false
		}
	}
	return true
}

// 暴力解烦 递归 时O(n)空O(n) 不超过32位 超过了得换成类型
func nQueen2(n int) int {
	if n < 1 || n > 32 {
		return 0
	}

	var limit int
	if n == 32 {
		limit = -1
	} else {
		limit = 1<<n - 1
	}

	return process2(limit, 0, 0, 0)
}

// 基于位运算的解法，利用位运算的高效性来减少递归的复杂度
// colLim 列的限制,1的位置不能放皇后,0的位置可以
// leftDiaLim 左斜线的限制
// rightDiaLim 右斜线的限制
func process2(limit, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == limit {
		return 1
	}
	res, mostRightOne := 0, 0
	// 所有候选皇后的位置
	pos := limit & ^(colLim | leftDiaLim | rightDiaLim)
	for pos != 0 {
		mostRightOne = pos & (^pos + 1) // 最右侧的1
		pos = pos - mostRightOne
		res += process2(limit, colLim|mostRightOne,
			(leftDiaLim|mostRightOne)<<1, (rightDiaLim|mostRightOne)>>1)
	}
	return res
}

func main() {
	testN := 14
	results := make(chan string, 2)

	// 测试 nQueen1
	go func() {
		start := time.Now()
		res1 := nQueen1(testN)
		elapsed1 := time.Since(start)
		results <- fmt.Sprintf("nQueen1(%d) = %d, Time: %v\n", testN, res1, elapsed1)
	}()

	// 测试 nQueen2
	go func() {
		start := time.Now()
		res2 := nQueen2(testN)
		elapsed2 := time.Since(start)
		results <- fmt.Sprintf("nQueen2(%d) = %d, Time: %v\n", testN, res2, elapsed2)
	}()

	// 打印结果
	for i := 0; i < 2; i++ {
		fmt.Print(<-results)
	}
}
