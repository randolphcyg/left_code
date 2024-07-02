package main

import (
	"fmt"
	"math"

	"left_code/sort/helper"
)

// selectionSort 选择排序 不稳定 时O(N^2) 空O(1)
func selectionSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ { // i~ N-1
		minIndex := i
		for j := i + 1; j < len(arr); j++ { // i~ N-1上找最小值的下标
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

// bubbleSort 冒泡排序 稳定 时O(N^2) 空O(1)
func bubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := len(arr) - 1; i > 0; i-- { // 0 ~ e
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// insertSort 插入排序 稳定 时O(N^2) 空O(1)
func insertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ { // 0 ~ e // 从第二个元素开始
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- { // 比较当前元素与前一个元素，如果当前元素较小则交换
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}

// int数组中只有一种数 出现了奇数次，怎么找到这1个数？
func printOddTimesOneNum(arr []int) int {
	eor := 0
	for _, cur := range arr {
		eor ^= cur
	}
	return eor
}

// int数组中只有两种数 出现了奇数次，怎么找到这2个数？
// 异或运算：异或运算具有交换律和结合律，相同的数异或结果为 0，不同的数异或结果为 1。
func printOddTimesTwoNum(arr []int) (num1, num2 int) {
	eor := 0
	for _, item := range arr {
		eor ^= item
	}

	// 最低位的 1：用于找到 eor 的最低位的 1，-eor 是 eor 的补码表示。
	rightOne := eor & -eor // 得到最低位的1 这个位是 eor 中两个不同奇数的二进制表示中不同的一位。
	onlyOne := 0
	// 利用 rightOne 将数组分成两部分，一个部分包含 rightOne 位为 1 的数，另一个部分包含 rightOne 位为 0 的数。
	// 然后分别对这两部分进行异或运算，得到这两个不同的奇数。
	for _, cur := range arr {
		if cur&rightOne != 0 {
			onlyOne ^= cur
		}
	}

	return onlyOne, eor ^ onlyOne
}

// 二分查找 在一个有序数组中，找某个数是否存在
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // 未找到目标值
}

// 在一个有序数组中，找>=某个数最左侧的位置
func findFirstGreaterOrEqual(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] >= target {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

// 局部最小问题
func findLocalMinima(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	N := len(arr)
	if N == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[N-1] < arr[N-2] {
		return N - 1
	}

	left := 1
	right := N - 2
	mid := 0
	for left < right {
		mid = left + ((right - left) >> 1)
		if arr[mid-1] < arr[mid] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}

// 递归获取数组最大值
// T(N)=2*T(N/2)+O(1)	Master公式：a=2 b=2 d=0 ∵ log(b,a) > d ∴ O(N*log(b,a))
func getMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, L int, R int) int {
	if L == R { // L..R范围上只有一个数，直接返回
		return arr[L]
	}
	mid := L + (R-L)>>1 // 中点 不会溢出且更快
	leftMax := process(arr, L, mid)
	rightMax := process(arr, mid+1, R)
	return int(math.Max(float64(leftMax), float64(rightMax)))
}

func main() {
	arraySize := 100
	maxValue := 1000
	testCount := 5000

	helper.TestSortingConsistency(bubbleSort, arraySize, maxValue, testCount)
	helper.TestSortingConsistency(selectionSort, arraySize, maxValue, testCount)
	helper.TestSortingConsistency(insertSort, arraySize, maxValue, testCount)

	array := []int{8, 0, 6, 5, 200, 2, 5, 10}
	fmt.Println("source array: ", array)
	bubbleSortRes := bubbleSort(array)
	fmt.Println("bubble sort: ", bubbleSortRes)

	selectionSortRes := selectionSort(array)
	fmt.Println("selection sort: ", selectionSortRes)

	sinsertSortRes := insertSort(array)
	fmt.Println("selection sort: ", sinsertSortRes)

	// 测试用例
	arr := []int{1, 2, 3, 2, 1}
	fmt.Println(printOddTimesOneNum(arr)) // 输出: 3

	arr = []int{1, 2, 3, 2, 1, 4}
	fmt.Println(printOddTimesTwoNum(arr)) // 输出: 3 4

	arr = []int{5, 7, 5, 7, 5, 9, 9, 3}
	fmt.Println(printOddTimesTwoNum(arr)) // 输出: 3 5

	arr = []int{10, 14, 14, 10, 15, 15, 20, 25}
	fmt.Println(printOddTimesTwoNum(arr)) // 输出: 20 25

	// 二分查找 测试用例
	index := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	if index != -1 {
		fmt.Printf("元素 %d 在数组中的索引是 %d\n", 5, index)
	} else {
		fmt.Printf("数组中没有找到元素 %d\n", 5)
	}

	fmt.Println(findFirstGreaterOrEqual([]int{1, 3, 5, 7, 9}, 6))  // 输出: 3
	fmt.Println(findFirstGreaterOrEqual([]int{2, 4, 6, 8, 10}, 5)) // 输出: 2

	// 找到一个局部最小即可
	arr = []int{9, 6, 3, 2, 8, 14, 1, 5, 7}
	idx := findLocalMinima(arr)
	fmt.Println(arr, "局部最小数：", arr[idx]) // Output: 8

	fmt.Println(getMax([]int{9, 6, 3, 2, 8, 14, 1, 5, 7}))
}
