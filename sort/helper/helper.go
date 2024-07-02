package helper

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"slices"
	"time"
)

/*
对数器
*/

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

// SortFunc 排序函数类型
type SortFunc func([]int) []int

// TestSortingConsistency 对数器-测试排序函数
func TestSortingConsistency(sortFunc SortFunc, arraySize int, maxValue int, testCount int) {
	funcName := runtime.FuncForPC(reflect.ValueOf(sortFunc).Pointer()).Name()
	for i := 0; i < testCount; i++ {
		originalArray := generateRandomArray(arraySize, maxValue)
		customSortedArray := make([]int, len(originalArray))
		officialSortedArray := make([]int, len(originalArray))

		copy(customSortedArray, originalArray)
		copy(officialSortedArray, originalArray)

		sortFunc(customSortedArray)
		slices.Sort(officialSortedArray) // 官方排序函数

		if !reflect.DeepEqual(customSortedArray, officialSortedArray) {
			fmt.Println("异常点", i+1)
			fmt.Println("原数组: ", originalArray)
			fmt.Println("自定义排序结果: ", customSortedArray)
			fmt.Println("官方库排序结果: ", officialSortedArray)
			fmt.Println(funcName, ":测试失败！！！")
			return
		}
	}
	fmt.Println()
	fmt.Println(funcName, ":测试通过~")
	return
}

// GreedyFunc 贪心策略函数类型
type GreedyFunc func([]int) int

// TestGreedyConsistency 对数器-测试贪心策略函数
func TestGreedyConsistency(greedyFunc1, greedyFunc2 GreedyFunc, arraySize int, maxValue int, testCount int, debug bool) {
	funcName1 := runtime.FuncForPC(reflect.ValueOf(greedyFunc1).Pointer()).Name()
	funcName2 := runtime.FuncForPC(reflect.ValueOf(greedyFunc2).Pointer()).Name()

	for i := 0; i < testCount; i++ {
		originalArray := generateRandomArray(arraySize, maxValue)
		res1 := greedyFunc1(originalArray)
		res2 := greedyFunc2(originalArray)

		if res1 != res2 {
			fmt.Println("异常点", i+1)
			fmt.Println("原数组: ", originalArray)
			fmt.Println("贪心策略函数1结果: ", res1)
			fmt.Println("暴力函数2结果: ", res2)
			fmt.Println(funcName1, "和", funcName2, ":测试失败！！！")
			return
		}

		if debug {
			fmt.Println("随机数组：", originalArray)
			fmt.Println("结果：", res1)
		}

		// 计算并显示进度百分比
		progress := float64(i+1) / float64(testCount) * 100
		fmt.Printf("\r进度: %.2f%%", progress)
	}
	fmt.Println()
	fmt.Println(funcName1, "和", funcName2, ":测试通过~")
	return
}
