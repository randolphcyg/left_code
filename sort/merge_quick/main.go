package main

import (
	"math/rand"

	"left_code/sort/helper"
)

// mergeSort 归并排序 稳定 时O(N*logN) 空O(N)
// T(N) = 2*T(N/2) + O(N)	Master公式：a=2 b=2 d=1	∵ log(b,a) = d ∴ O(N*logN)
func mergeSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	process(arr, 0, len(arr)-1)
	return arr
}

func process(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	process(arr, l, mid)
	process(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) {
	help := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := m + 1

	for p1 <= m && p2 <= r {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	// 复制左边剩余的元素
	for p1 <= m {
		help[i] = arr[p1]
		i++
		p1++
	}
	// 复制右边剩余的元素
	for p2 <= r {
		help[i] = arr[p2]
		i++
		p2++
	}

	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}
}

// smallSum 小和问题
func smallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return processSmallSum(arr, 0, len(arr)-1)
}

func processSmallSum(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)>>1
	return processSmallSum(arr, l, mid) + processSmallSum(arr, mid+1, r) + mergeSmallSum(arr, l, mid, r)
}

func mergeSmallSum(arr []int, l, m, r int) int {
	help := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := m + 1
	res := 0

	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			res += (r - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	// 复制左边剩余的元素
	for p1 <= m {
		help[i] = arr[p1]
		i++
		p1++
	}
	// 复制右边剩余的元素
	for p2 <= r {
		help[i] = arr[p2]
		i++
		p2++
	}

	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}
	return res
}

// 快排 不稳定 时O(N*logN) 空O(logN)
func quickSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	quickSortProcess(arr, 0, len(arr)-1)
	return arr
}

func quickSortProcess(arr []int, l, r int) {
	if l < r {
		randomIndex := l + rand.Intn(r-l+1)
		arr[randomIndex], arr[r] = arr[r], arr[randomIndex]
		p := partition(arr, l, r)
		quickSortProcess(arr, l, p[0]-1) // <区
		quickSortProcess(arr, p[1]+1, r) // >区
	}
}

/*
处理arr[l..r]的函数
默认以arr[r]做划分，arr[r] 	->		p	<p	==p	>p
返回等于区域(左边界，右边界)。所以返回一个长度为2的数组
*/
func partition(arr []int, l, r int) []int {
	less := l - 1 // <区右边界
	more := r     // >区左边界
	cur := l      // 当前数的位置

	for cur < more { // cur表示当前数的位置 arr[r] -> 划分值
		if arr[cur] < arr[r] { // 当前数 < 划分值
			less++
			arr[less], arr[cur] = arr[cur], arr[less]
			cur++
		} else if arr[cur] > arr[r] { // 当前数 > 划分值
			more--
			arr[more], arr[cur] = arr[cur], arr[more]
		} else { // 当前数 == 划分值
			cur++
		}
	}
	arr[more], arr[r] = arr[r], arr[more]
	return []int{less + 1, more}
}

func main() {
	//arr := []int{3, 2, 1, 5, 6, 4}
	//res := quickSort(arr)
	//fmt.Println(res)

	arraySize := 100
	maxValue := 1000
	testCount := 5000 // 测试5千次

	helper.TestSortingConsistency(mergeSort, arraySize, maxValue, testCount)
	helper.TestSortingConsistency(quickSort, arraySize, maxValue, testCount)

	//arr := []int{9, 6, 3, 2, 8, 14, 1, 5, 7}
	//fmt.Println(arr)
	//res := mergeSort(arr)
	//fmt.Println(res)

	//res := smallSum([]int{1, 3, 4, 2, 5})
	//fmt.Println(res)
}
