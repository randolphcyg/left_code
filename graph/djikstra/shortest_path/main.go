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
		from, to, weight := link[1], link[2], link[3]
		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		edge := NewEdge(weight, fromNode, toNode)
		fromNode.Edges = append(fromNode.Edges, edge)
		fromNode.Nexts = append(fromNode.Nexts, toNode)
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
	heap.Push(pq, NewEdge(0, start, start))

	for !pq.IsEmpty() {
		edge, _ := heap.Pop(pq).(*Edge)
		node := edge.To
		distance := edge.Weight

		for _, e := range node.Edges {
			toNode := e.To
			if toNode == node {
				toNode = e.From
			}
			newDist := distance + e.Weight
			if e.From != start { // 只有在换乘时才应用优惠
				newDist -= 1
			}
			if newDist < distanceMap[toNode] {
				distanceMap[toNode] = newDist
				heap.Push(pq, NewEdge(newDist, node, toNode))
			}
		}
	}
	return distanceMap
}

// calculateMinCost 计算最小路径
func calculateMinCost(distanceMap map[*Node]int, dstNode *Node) string {
	if distanceMap[dstNode] == math.MaxInt64 {
		return "NA"
	}
	return fmt.Sprintf("%d", distanceMap[dstNode])
}

/*
输入：
4 5
1 1 2 3
2 1 3 3
3 1 4 4
4 2 3 5
5 3 4 3
1 3
输出
3
*/
func main() {
	// 示例输入
	var N, M int // N个地点 M个公交路线
	fmt.Scan(&N, &M)
	links := make([][]int, M)
	for i := 0; i < M; i++ {
		links[i] = make([]int, 4)
		fmt.Scan(&links[i][0], &links[i][1], &links[i][2], &links[i][3])
	}
	var srcNodeIndex, dstNodeIndex int // 出发地和目的地
	fmt.Scan(&srcNodeIndex, &dstNodeIndex)

	// 创建图
	graph := createGraph(N, links)

	// 获取起始节点和目标节点
	srcNode, dstNode := graph.Nodes[srcNodeIndex], graph.Nodes[dstNodeIndex]

	// 运行 Dijkstra 算法
	result := dijkstra(graph, srcNode)

	// 计算最便宜的换乘路线
	minCost := calculateMinCost(result, dstNode)

	// 打印结果
	fmt.Println(minCost)
}
