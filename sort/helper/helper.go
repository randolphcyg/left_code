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
测试方法
*/

// 生成随机数组
func generateRandomArray(size int, maxValue int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.Intn(maxValue)
	}
	return arr
}

// SortFunc 排序函数类型
type SortFunc func([]int) []int

// TestSortingConsistency 测试函数
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
	fmt.Println(funcName, ":测试通过~")
	return
}

/*
比较器
*/

func Comparator(arg0, arg1 int) int {
	return arg1 - arg0
}
