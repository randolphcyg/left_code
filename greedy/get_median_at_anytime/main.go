package main

import (
	"container/heap"
	"fmt"
)

// MaxHeap 大根堆
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinHeap 小根堆
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MedianFinder 数据流中位数查找器
type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

// NewMedianFinder 初始化
func NewMedianFinder() *MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return &MedianFinder{maxHeap, minHeap}
}

// AddNum 添加一个数到数据流
func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.maxHeap, num)
	heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))

	if mf.maxHeap.Len() < mf.minHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

// FindMedian 获取当前数据流的中位数
func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	}
	return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
}

func main() {
	mf := NewMedianFinder()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, num := range nums {
		mf.AddNum(num)
		fmt.Printf("Added: %d, Current Median: %.1f\n", num, mf.FindMedian())
	}
}
