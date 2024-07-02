package main

import (
	"container/heap"
	"fmt"
	"math"

	"left_code/sort/helper"
)

/*暴力求解*/

// bruteForceSplitGold 使用暴力递归解决分金问题
func bruteForceSplitGold(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return process(arr, 0)
}

// process 是实际的递归函数，用于计算最小花费
func process(arr []int, preCost int) int {
	if len(arr) == 1 {
		return preCost
	}

	minCost := math.MaxInt32

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			newArr := mergeTwoElements(arr, i, j)
			curCost := preCost + arr[i] + arr[j]
			minCost = int(math.Min(float64(minCost), float64(process(newArr, curCost))))
		}
	}

	return minCost
}

// mergeTwoElements 合并数组中的两个元素
func mergeTwoElements(arr []int, i, j int) []int {
	newArr := make([]int, 0, len(arr)-1)
	for k := 0; k < len(arr); k++ {
		if k != i && k != j {
			newArr = append(newArr, arr[k])
		}
	}
	newArr = append(newArr, arr[i]+arr[j])
	return newArr
}

/*优先队列*/

// PriorityQueue 优先队列 java/go底层就是小根堆
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

// 分金问题 (哈夫曼编码)
func lessMoneySplitGold(arr []int) int {
	pQ := NewPriorityQueue()
	for _, item := range arr {
		heap.Push(pQ, item)
	}
	sum, cur := 0, 0
	for pQ.Len() > 1 {
		cur = heap.Pop(pQ).(int) + heap.Pop(pQ).(int) //每次取出2两个相加，将和放回
		sum += cur
		heap.Push(pQ, cur)
	}
	return sum
}

func main() {
	// 测试案例
	arr := []int{10, 20, 30}
	result := lessMoneySplitGold(arr)
	fmt.Printf("最小花费为: %d\n", result)

	result = bruteForceSplitGold(arr)
	fmt.Printf("暴力求解最小花费为: %d\n", result)
	fmt.Println()

	arraySize := 8
	maxValue := 20
	testCount := 10
	helper.TestGreedyConsistency(lessMoneySplitGold, bruteForceSplitGold, arraySize, maxValue, testCount, false)
}
