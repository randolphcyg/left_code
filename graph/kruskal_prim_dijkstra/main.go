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
	/*
		4 2 2
		1 2 3
		3 2 5
	*/
	for edge, _ := range resKruskal {
		fmt.Println(edge.From.Val, edge.To.Val, edge.Weight)
	}

	fmt.Println("prim: ")
	resPrim := prim(graph)
	/*
		4 2 2
		2 1 3
		2 3 5
	*/
	for edge, _ := range resPrim {
		fmt.Println(edge.From.Val, edge.To.Val, edge.Weight)
	}

	fmt.Println("dijkstra: ")
	resDijkstra := dijkstra(graph.Nodes[1])
	/*
		4 5 节点1到节点4距离最短距离是5
		1 0
		2 3
		3 8 节点1到节点3距离最短距离是8
	*/
	for node, distance := range resDijkstra {
		fmt.Println(node.Val, distance)
	}

	// 运行dijkstra算法
	fmt.Println("dijkstra2: ")
	resDijkstra2 := dijkstra2(graph.Nodes[1], len(graph.Nodes))
	/*
		4 5
		3 8
		1 0
		2 3
	*/
	for node, distance := range resDijkstra2 {
		fmt.Println(node.Val, distance)
	}
}
