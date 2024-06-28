package main

import "fmt"

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

// Queue 定义一个队列结构体
type Queue struct {
	elements []*Node
}

// NewQueue 初始化并返回一个新的 Queue 实例
func NewQueue() *Queue {
	return &Queue{
		elements: []*Node{},
	}
}

// Enqueue 将一个节点入队
func (q *Queue) Enqueue(node *Node) {
	q.elements = append(q.elements, node)
}

// Dequeue 将一个节点出队并返回
func (q *Queue) Dequeue() (*Node, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	node := q.elements[0]
	q.elements = q.elements[1:]
	return node, true
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size 返回队列的大小
func (q *Queue) Size() int {
	return len(q.elements)
}

func sortedTopology(graph *Graph) []*Node {
	// key:某个node value:剩余的入度
	inMap := map[*Node]int{}
	// 入度为0的点，才能进这个队列
	zeroInQueue := NewQueue()
	for _, node := range graph.Nodes {
		inMap[node] = node.In
		if node.In == 0 {
			zeroInQueue.Enqueue(node)
		}
	}
	// 拓扑排序的结果，依次加入result
	var result []*Node
	for !zeroInQueue.IsEmpty() {
		cur, _ := zeroInQueue.Dequeue()
		result = append(result, cur)
		for _, next := range cur.Nexts {
			inMap[next] = inMap[next] - 1
			if inMap[next] == 0 {
				zeroInQueue.Enqueue(next)
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
