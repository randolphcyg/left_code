package main

import (
	"math"

	"sort/helper"
)

func radixSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	radixSortProcess(arr, 0, len(arr)-1, maxbits(arr))
	return arr
}

func maxbits(arr []int) int {
	max := math.MaxInt
	for i := 0; i < len(arr); i++ {
		max = int(math.Max(float64(max), float64(arr[i])))
	}
	res := 0
	for max != 0 {
		res++
		max /= 10
	}
	return res
}

// getDigit 获取数字 num 从右到左第 pos 位的数字
func getDigit(num, pos int) int {
	return (num / int(math.Pow10(pos))) % 10
}

func radixSortProcess(arr []int, l, r, digit int) []int {
	radix := 10
	// 有多少个数准备多少个辅助空间
	bucket := make([]int, r-l+1)
	for d := 0; d <= digit; d++ { // 有多少位就进出几次
		// 10个空间
		// count[0] 当前位(d位)是0的数字有多少个
		// count[1] 当前位(d位)是0和1的数字有多少个
		// count[2] 当前位(d位)是0、1、2的数字有多少个
		// count[i] 当前位(d位)是0~i的数字有多少个
		count := make([]int, radix)
		for i := l; i <= r; i++ {
			j := getDigit(arr[i], d)
			count[j]++
		}
		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		for i := r; i >= l; i-- {
			j := getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}
		for i, j := l, 0; i <= r; i, j = i+1, j+1 {
			arr[i] = bucket[j]
		}
	}
	return arr
}

func main() {
	arraySize := 100
	maxValue := 1000
	testCount := 5000 // 测试5千次
	helper.TestSortingConsistency(radixSort, arraySize, maxValue, testCount)
}
