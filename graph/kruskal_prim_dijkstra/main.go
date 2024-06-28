package main

import (
	"container/heap"
	"fmt"
	"math"
)

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

// createGraph 根据给定的二维数组填充图
func createGraph(matrix [][]int) *Graph {
	graph := NewGraph()
	for _, row := range matrix {
		fromVal, toVal, weight := row[0], row[1], row[2]

		if _, exists := graph.Nodes[fromVal]; !exists {
			graph.Nodes[fromVal] = &Node{Val: fromVal}
		}
		if _, exists := graph.Nodes[toVal]; !exists {
			graph.Nodes[toVal] = &Node{Val: toVal}
		}

		fromNode := graph.Nodes[fromVal]
		toNode := graph.Nodes[toVal]

		edge := &Edge{
			Weight: weight,
			From:   fromNode,
			To:     toNode,
		}

		fromNode.Nexts = append(fromNode.Nexts, toNode)
		fromNode.Edges = append(fromNode.Edges, edge)
		fromNode.Out++
		toNode.In++

		graph.Edges[edge] = struct{}{}
	}
	return graph
}

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

// MySets 我的集合们
type MySets struct {
	setMap map[*Node][]*Node
}

func NewMySets(nodes []*Node) *MySets {
	setMap := make(map[*Node][]*Node)
	for _, cur := range nodes {
		var set []*Node
		set = append(set, cur)
		setMap[cur] = set
	}

	return &MySets{setMap: setMap}
}

func (m *MySets) isSameSet(from, to *Node) bool {
	fromSet := m.setMap[from]
	toSet := m.setMap[to]
	return fromSet[0] == toSet[0]
}

func (m *MySets) union(from, to *Node) {
	fromSet := m.setMap[from]
	toSet := m.setMap[to]

	for _, toNode := range toSet {
		fromSet = append(fromSet, toNode)
		m.setMap[toNode] = fromSet
	}
}

func kruskal(graph *Graph) map[*Edge]struct{} {
	var nodes []*Node
	for _, node := range graph.Nodes {
		nodes = append(nodes, node)
	}
	mySets := NewMySets(nodes)
	priorityQueue := NewPriorityQueue()
	for edge, _ := range graph.Edges {
		priorityQueue.Push(edge)
	}
	result := make(map[*Edge]struct{})
	for !priorityQueue.IsEmpty() {
		edge := priorityQueue.Pop().(*Edge)
		if !mySets.isSameSet(edge.From, edge.To) {
			result[edge] = struct{}{}
			mySets.union(edge.From, edge.To)
		}
	}

	return result
}

func prim(graph *Graph) map[*Edge]struct{} {
	priorityQueue := NewPriorityQueue()
	set := map[*Node]struct{}{}
	result := map[*Edge]struct{}{} // 依次挑选的边在result里

	for _, node := range graph.Nodes { // 随便挑选一个点
		if _, ok := set[node]; !ok {
			set[node] = struct{}{}
			for _, edge := range node.Edges { // 由一个点，解锁所有相连的边
				priorityQueue.Push(edge)
			}

			for !priorityQueue.IsEmpty() {
				edge := priorityQueue.Pop().(*Edge) // 弹出解锁的边中，最小的边
				toNode := edge.To                   // 可能的一个新的点
				if _, ok := set[toNode]; !ok {      //不含有的时候，就是新的点
					set[toNode] = struct{}{}
					result[edge] = struct{}{}
					for _, edge := range toNode.Edges {
						priorityQueue.Push(edge)
					}
				}
			}
		}
	}

	return result
}

func dijkstra(head *Node) map[*Node]int {
	// 从head出发到所有点的最小距离
	// key：从head出发到达key
	// value：从head出发到达key的最小距离
	// 如果在表中，没有T的记录，含义是从head出发到T这个点的距离为正无穷
	distanceMap := map[*Node]int{}
	distanceMap[head] = 0
	// 已经求过距离的节点，存在selectedNodes中，以后再也不碰
	selectedNodes := map[*Node]struct{}{}
	minNode := getMinDistanceAndUnselectedNode(distanceMap, selectedNodes)

	for minNode != nil {
		distance := distanceMap[minNode]
		for _, edge := range minNode.Edges {
			toNode := edge.To
			if _, ok := distanceMap[toNode]; !ok {
				distanceMap[toNode] = distance + edge.Weight
			}
			distanceMap[toNode] = int(math.Min(float64(distanceMap[toNode]), float64(distance+edge.Weight)))
		}
		selectedNodes[minNode] = struct{}{}
		minNode = getMinDistanceAndUnselectedNode(distanceMap, selectedNodes)
	}
	return distanceMap
}

func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, touchedNodes map[*Node]struct{}) (minNode *Node) {
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		if _, ok := touchedNodes[node]; !ok && distance < minDistance {
			minNode = node
			minDistance = distance
		}
	}
	return minNode
}

func main() {
	// 无向图
	matrix := [][]int{
		{1, 2, 3},
		{1, 3, 100},
		{1, 4, 7},
		{2, 3, 5},
		{2, 4, 2},
		{3, 4, 1000},

		{2, 1, 3},
		{3, 1, 100},
		{4, 1, 7},
		{3, 2, 5},
		{4, 2, 2},
		{4, 3, 1000},
	}

	graph := createGraph(matrix)

	fmt.Println("kruskal: ")
	resKruskal := kruskal(graph)
	for edge, _ := range resKruskal {
		fmt.Println(edge.From.Val, edge.To.Val, edge.Weight)
	}

	fmt.Println("prim: ")
	resPrim := prim(graph)
	for edge, _ := range resPrim {
		fmt.Println(edge.From.Val, edge.To.Val, edge.Weight)
	}

	fmt.Println("dijkstra: ")
	resDjikstra := dijkstra(graph.Nodes[1])
	for node, distance := range resDjikstra {
		fmt.Println(node.Val, distance)
	}
}
