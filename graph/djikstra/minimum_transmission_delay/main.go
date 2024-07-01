package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*自定义图结构*/

type Node struct {
	Val   int
	In    int
	Out   int
	Nexts []*Node
	Edges []*Edge
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
	Weight int
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

func createGraph(N int, links [][]int) *Graph {
	graph := NewGraph()
	for i := 1; i <= N; i++ {
		graph.Nodes[i] = NewNode(i)
	}
	for _, link := range links {
		from, to, weight := link[0], link[1], link[2]
		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		edge := NewEdge(weight, fromNode, toNode)
		fromNode.Edges = append(fromNode.Edges, edge)
		toNode.Edges = append(toNode.Edges, edge)
		graph.Edges[edge] = struct{}{}
	}
	return graph
}

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

// Dijkstra 算法实现 函数计算从起始结点到所有其他结点的最短路径
func dijkstra(graph *Graph, start *Node) map[*Node]int {
	distanceMap := make(map[*Node]int)
	for _, node := range graph.Nodes {
		distanceMap[node] = math.MaxInt64
	}
	distanceMap[start] = 0

	pq := NewPriorityQueue()
	pq.Push(NewEdge(0, start, start))

	for !pq.IsEmpty() {
		edge := pq.Pop().(*Edge)
		node := edge.To
		distance := edge.Weight

		for _, e := range node.Edges {
			toNode := e.To
			newDist := distance + e.Weight
			if newDist < distanceMap[toNode] {
				distanceMap[toNode] = newDist
				pq.Push(NewEdge(newDist, node, toNode))
			}
		}
	}
	return distanceMap
}

// 计算最小传输时延
func calculateMinDelay(distanceMap map[*Node]int, dstNode *Node) int {
	if distanceMap[dstNode] == math.MaxInt64 {
		return -1
	}
	return distanceMap[dstNode]
}

/*
输入
3 3
1 2 11
2 3 13
1 3 50
1 3
输出
24
*/
func main() {
	// 示例输入
	var N, M int // 网络结点的个数N以及时延列表长度M
	fmt.Scan(&N, &M)
	links := make([][]int, M)
	for i := 0; i < M; i++ {
		links[i] = make([]int, 3)
		fmt.Scan(&links[i][0], &links[i][1], &links[i][2])
	}
	var srcNodeIndex, dstNodeIndex int // 源结点和目的结点
	fmt.Scan(&srcNodeIndex, &dstNodeIndex)

	// 创建图
	graph := createGraph(N, links)

	// 获取起始节点和目标节点
	srcNode, dstNode := graph.Nodes[srcNodeIndex], graph.Nodes[dstNodeIndex]

	// 运行 Dijkstra 算法
	result := dijkstra(graph, srcNode)

	// 计算最小传输时延
	minDelay := calculateMinDelay(result, dstNode)

	// 打印结果
	fmt.Println(minDelay)
}
