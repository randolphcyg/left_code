package main

import (
	"fmt"
)

// 双端队列 解决滑动窗口问题
func maxSlidingWindow(arr []int, w int) []int {
	if len(arr) == 0 || w == 0 {
		return []int{}
	}

	n := len(arr)
	res := make([]int, n-w+1)
	deque := []int{}

	for i := 0; i < n; i++ {
		// 如果队列中的第一个元素不在当前窗口范围内，则移除
		if len(deque) > 0 && deque[0] < i-w+1 {
			deque = deque[1:]
		}

		// 移除队列中所有小于当前元素的索引
		for len(deque) > 0 && arr[deque[len(deque)-1]] <= arr[i] {
			deque = deque[:len(deque)-1]
		}

		// 添加当前元素的索引到双端队列中
		deque = append(deque, i)

		// 如果当前索引大于等于 w-1，记录当前窗口的最大值
		if i >= w-1 {
			res[i-w+1] = arr[deque[0]]
		}
	}

	return res
}

func main() {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	w := 3
	result := maxSlidingWindow(arr, w)
	fmt.Println(result) // 输出: [5 5 5 4 6 7]
}
