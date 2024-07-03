package main

import "container/heap"

/*优先队列*/

// PriorityQueue 实现优先队列
type PriorityQueue []*Edge

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Weight < (*pq)[j].Weight
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

// prim 算法实现
func prim(graph *Graph) map[*Edge]struct{} {
	priorityQueue := NewPriorityQueue()
	set := map[*Node]struct{}{}
	result := map[*Edge]struct{}{} // 依次挑选的边在result里

	for _, node := range graph.Nodes { // 随便挑选一个点
		if _, ok := set[node]; !ok {
			set[node] = struct{}{}
			for _, edge := range node.Edges { // 由一个点，解锁所有相连的边
				heap.Push(priorityQueue, edge)
			}

			for !priorityQueue.IsEmpty() {
				edge := heap.Pop(priorityQueue).(*Edge) // 弹出解锁的边中，最小的边
				toNode := edge.To                       // 可能的一个新的点
				if _, ok := set[toNode]; !ok {          //不含有的时候，就是新的点
					set[toNode] = struct{}{}
					result[edge] = struct{}{}
					for _, edge := range toNode.Edges {
						heap.Push(priorityQueue, edge)
					}
				}
			}
		}
	}

	return result
}
