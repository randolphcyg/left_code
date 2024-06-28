package main

import (
	"container/heap"
	"fmt"

	"sort/helper"
)

// 某个数现在处在index位置，往上继续移动
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

// heapify
func heapify(arr []int, index, heapSize int) {
	left := 2*index + 1   // 左孩子下标
	for left < heapSize { // 下方还有孩子的时候
		largest := left
		// 俩孩子中，谁的值大，把下标给largest
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}

		// 父和孩子之间，谁的值大下标就给largest
		if arr[largest] <= arr[index] {
			break
		}

		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2*index + 1
	}
}

// heapSort 堆排 O(N*logN) 空间O(1)
func heapSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	// 构建大根堆
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}

	// 只让数组生成大根堆 用这个更快
	//for i := len(arr); i >= 0; i-- {
	//	heapify(arr, i, len(arr))
	//}

	heapSize := len(arr)
	heapSize--
	arr[0], arr[heapSize] = arr[heapSize], arr[0]

	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize--
		arr[0], arr[heapSize] = arr[heapSize], arr[0]
	}
	return arr
}

// 堆

// PriorityQueue 定义一个类型，它实现了 heap.Interface 接口
type PriorityQueue []int

func (pq *PriorityQueue) Len() int           { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i] < (*pq)[j] }
func (pq *PriorityQueue) Swap(i, j int)      { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

// 优先队列 java底层就是小根堆
func sortedArrDistanceLessK(arr []int, k int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	h := NewPriorityQueue()

	// 初始化小根堆
	for i := 0; i <= k && i < len(arr); i++ {
		heap.Push(h, arr[i])
	}

	result := make([]int, 0, len(arr))
	for i := k + 1; i < len(arr); i++ {
		heap.Push(h, arr[i])
		result = append(result, heap.Pop(h).(int))
	}
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(int))
	}

	return result
}

func main() {
	arraySize := 100
	maxValue := 1000
	testCount := 5000 // 测试5千次

	helper.TestSortingConsistency(heapSort, arraySize, maxValue, testCount)

	arr := []int{6, 5, 3, 2, 8, 10, 9}
	k := 3
	sortedArr := sortedArrDistanceLessK(arr, k)
	fmt.Println("Sorted array:", sortedArr)
}
