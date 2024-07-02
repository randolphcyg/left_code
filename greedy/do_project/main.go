package main

import (
	"container/heap"
	"fmt"
)

/*
TODO 本题不同点是可以同时支持大根堆 小根堆
*/

// Node 项目节点
type Node struct {
	P int // 利润
	C int // 投资
}

func NewNode(p, c int) *Node {
	return &Node{P: p, C: c}
}

/*优先队列*/

// PriorityQueue 优先队列
type PriorityQueue struct {
	items    []*Node
	lessFunc func(i, j *Node) bool
}

func (pq PriorityQueue) Len() int { return len(pq.items) }

func (pq PriorityQueue) Less(i, j int) bool { return pq.lessFunc(pq.items[i], pq.items[j]) }

func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.items = append(pq.items, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	x := old[n-1]
	pq.items = old[0 : n-1]
	return x
}

func NewPriorityQueue(lessFunc func(i, j *Node) bool) *PriorityQueue {
	pq := &PriorityQueue{
		items:    make([]*Node, 0),
		lessFunc: lessFunc,
	}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

// Peek 查看堆顶元素而不删除它
func (pq *PriorityQueue) Peek() *Node {
	if pq.Len() > 0 {
		return pq.items[0]
	}
	return nil
}

// 投资小根堆
func compareMin(i, j *Node) bool {
	return i.C < j.C
}

// 利润大根堆
func compareMax(i, j *Node) bool {
	return i.P > j.P
}

func findMaximizedCapital(k, W int, profits, capital []int) int {
	// minCostQ 是一个小根堆，用于根据项目的投资成本（C）排序
	minCostQ := NewPriorityQueue(compareMin)
	// maxProfitQ 是一个大根堆，用于根据项目的利润（P）排序
	maxProfitQ := NewPriorityQueue(compareMax)
	// 所有项目扔到被锁池中 花费组织的小根堆
	for i := 0; i < len(profits); i++ {
		heap.Push(minCostQ, NewNode(profits[i], capital[i]))
	}

	// 进行k轮选择，每轮选择能够进行的项目中利润最高的项目
	for i := 0; i < k; i++ {
		// 从minCostQ中弹出所有投资成本小于等于当前资本W的项目，并将他们压入maxProfitQ
		for !minCostQ.IsEmpty() && minCostQ.Peek().C <= W {
			heap.Push(maxProfitQ, heap.Pop(minCostQ).(*Node))
		}
		// 本轮没有可以进行的项目 返回当前的资本
		if maxProfitQ.IsEmpty() {
			return W
		}
		// 否则 从maxProfitQ弹出利润最高的项目 并将其利润加到当前总资本中
		W += heap.Pop(maxProfitQ).(*Node).P
	}

	return W
}

/*
TODO 暴力解法
*/

// 回溯法解决最大化资本问题
func findMaximizedCapitalBruteForce(k, W int, profits, capital []int) int {
	return helper(k, W, profits, capital, 0)
}

// helper 是一个辅助递归函数
func helper(k, W int, profits, capital []int, index int) int {
	if k == 0 || index == len(profits) {
		return W
	}

	maxCapital := W
	for i := index; i < len(profits); i++ {
		if capital[i] <= W {
			// 尝试选择第 i 个项目
			profit := profits[i]
			remainingProfits := append(append([]int{}, profits[:i]...), profits[i+1:]...)
			remainingCapitals := append(append([]int{}, capital[:i]...), capital[i+1:]...)
			maxCapital = max(maxCapital, helper(k-1, W+profit, remainingProfits, remainingCapitals, 0))
		}
	}
	return maxCapital
}

// 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	k := 4
	W := 6
	profits := []int{7, 5, 3, 2, 6}
	capital := []int{8, 7, 7, 3, 2}

	result1 := findMaximizedCapital(k, W, profits, capital)
	result2 := findMaximizedCapitalBruteForce(k, W, profits, capital)

	fmt.Printf("Test case:\n")
	fmt.Printf("k: %d, W: %d, profits: %v, capital: %v\n", k, W, profits, capital)
	fmt.Printf("Greedy method result: %d\n", result1)
	fmt.Printf("Brute force method result: %d\n", result2)

	if result1 != result2 {
		fmt.Println("Results do not match!")
	} else {
		fmt.Println("Results match!")
	}

	// 对数器测试
	TestGreedyConsistency(findMaximizedCapital, findMaximizedCapitalBruteForce, 8, 10, 1000, false)
}
