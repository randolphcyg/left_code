package main

import (
	"container/heap"
	"fmt"
)

/*自定义图结构*/

type Node struct {
	Val   int
	In    int     // 入度
	Out   int     // 出度
	Nexts []*Node // 发散出去的直接邻居点
	Edges []*Edge // 拥有的边
}

func NewNode(val int) *Node {
	return &Node{
		Val:   val,
		In:    0,
		Out:   0,
		Nexts: nil,
		Edges: nil,
	}
}

type Edge struct {
	Weight int // 权值 距离
	From   *Node
	To     *Node
}

func NewEdge(weight int, from, to *Node) *Edge {
	return &Edge{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

type Graph struct {
	Nodes map[int]*Node
	Edges map[*Edge]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
		Edges: make(map[*Edge]struct{}),
	}
}

// matrix N*3的矩阵
// 转换函数 二维数组转换为我们定义的数据结构
func createGraph(matrix [][]int) *Graph {
	graph := NewGraph()
	for i := 0; i < len(matrix); i++ {
		from, to, weight := matrix[i][0], matrix[i][1], matrix[i][2]
		if graph.Nodes[from] == nil {
			graph.Nodes[from] = NewNode(from)
		}
		if graph.Nodes[to] == nil {
			graph.Nodes[to] = NewNode(to)
		}
		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		newEdge := NewEdge(weight, fromNode, toNode)
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		fromNode.Out++
		toNode.In++
		fromNode.Edges = append(fromNode.Edges, newEdge)
		graph.Edges[newEdge] = struct{}{}
	}
	return graph
}

/*优先队列*/

// PriorityQueue 实现优先队列
type PriorityQueue []*Node

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Val < (*pq)[j].Val
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
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

// sortedTopology 拓扑排序
// 1.选一个没有前驱的顶点并输出
// 2.删除这个顶点及以它为起点的有向边
// 3.重复12步骤直到DAG为空或不存在无前驱的顶点
func sortedTopology(graph *Graph) []*Node {
	// key:某个node value:剩余的入度
	inMap := map[*Node]int{}
	// 入度为0的点，才能进这个队列
	zeroInQueue := NewPriorityQueue()
	for _, node := range graph.Nodes {
		inMap[node] = node.In
		if node.In == 0 {
			zeroInQueue.Push(node)
		}
	}
	// 拓扑排序的结果，依次加入result
	var result []*Node
	for !zeroInQueue.IsEmpty() {
		cur := zeroInQueue.Pop().(*Node)
		result = append(result, cur)
		for _, next := range cur.Nexts {
			inMap[next] = inMap[next] - 1
			if inMap[next] == 0 {
				zeroInQueue.Push(next)
			}
		}
	}
	return result
}

func main() {
	/*
			1
		   /|\
		  2 3 5
		  |X| |
		  4 6 7
	*/
	matrix := [][]int{
		{1, 2, 2},
		{1, 3, 1},
		{1, 5, 2},
		{2, 4, 1},
		{2, 6, 2},
		{3, 4, 2},
		{3, 6, 1},
		{5, 7, 1},
	}

	graph := createGraph(matrix)
	sortedTopology(graph)
	// 进行拓扑排序
	result := sortedTopology(graph)
	fmt.Println("Topological Sort Result:")
	for _, node := range result {
		fmt.Printf("Node %d\n", node.Val)
	}
}
