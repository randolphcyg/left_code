package main

import (
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

	fmt.Println("Nodes:")
	for _, node := range graph.Nodes {
		fmt.Printf("Node %d: In=%d, Out=%d\n", node.Val, node.In, node.Out)
	}

	fmt.Println("Edges:")
	for edge := range graph.Edges {
		fmt.Printf("Edge from %d to %d with weight %d\n", edge.From.Val, edge.To.Val, edge.Weight)
	}

	// 从节点1开始进行广度优先遍历
	startNode := graph.Nodes[1]
	fmt.Println("bfs:")
	bfs(startNode)
	fmt.Println("dfs:")
	dfs(startNode)
}
