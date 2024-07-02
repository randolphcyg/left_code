package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

// generateRandomArray 生成随机数组
func generateRandomArray(size int, maxVal int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.Intn(maxVal)
	}
	return arr
}

// generateRandomTestCase 生成随机测试用例
func generateRandomTestCase(size, maxValue int) (int, int, []int, []int) {
	k := rand.Intn(size) + 1
	W := rand.Intn(maxValue)
	profits := generateRandomArray(size, maxValue)
	capital := generateRandomArray(size, maxValue)
	return k, W, profits, capital
}

// TestGreedyConsistency 对数器-测试贪心策略函数
func TestGreedyConsistency(greedyFunc1, greedyFunc2 func(int, int, []int, []int) int, size, maxValue, testCount int, debug bool) {
	funcName1 := runtime.FuncForPC(reflect.ValueOf(greedyFunc1).Pointer()).Name()
	funcName2 := runtime.FuncForPC(reflect.ValueOf(greedyFunc2).Pointer()).Name()

	for i := 0; i < testCount; i++ {
		k, W, profits, capital := generateRandomTestCase(size, maxValue)
		res1 := greedyFunc1(k, W, profits, capital)
		res2 := greedyFunc2(k, W, profits, capital)

		if res1 != res2 {
			fmt.Println("异常点", i+1)
			fmt.Printf("k: %d, W: %d, profits: %v, capital: %v\n", k, W, profits, capital)
			fmt.Println("贪心策略函数1结果: ", res1)
			fmt.Println("暴力函数2结果: ", res2)
			fmt.Println(funcName1, "和", funcName2, ":测试失败！！！")
			return
		}

		if debug {
			fmt.Printf("测试用例 %d:\n", i+1)
			fmt.Printf("k: %d, W: %d, profits: %v, capital: %v\n", k, W, profits, capital)
			fmt.Println("结果：", res1)
		}

		progress := float64(i+1) / float64(testCount) * 100
		fmt.Printf("\r进度: %.2f%%", progress)
	}
	fmt.Println()
	fmt.Println(funcName1, "和", funcName2, ":测试通过~")
}
