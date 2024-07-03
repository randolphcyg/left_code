package main

import "container/heap"

// DisjointSet 并查集
type DisjointSet struct {
	parent map[*Node]*Node
	rank   map[*Node]int
}

func NewDisjointSet(nodes []*Node) *DisjointSet {
	ds := &DisjointSet{
		parent: make(map[*Node]*Node),
		rank:   make(map[*Node]int),
	}
	for _, node := range nodes {
		ds.parent[node] = node
		ds.rank[node] = 0
	}
	return ds
}

func (ds *DisjointSet) Find(node *Node) *Node {
	if ds.parent[node] != node {
		ds.parent[node] = ds.Find(ds.parent[node])
	}
	return ds.parent[node]
}

func (ds *DisjointSet) Union(node1, node2 *Node) {
	root1 := ds.Find(node1)
	root2 := ds.Find(node2)
	if root1 != root2 {
		if ds.rank[root1] > ds.rank[root2] {
			ds.parent[root2] = root1
		} else if ds.rank[root1] < ds.rank[root2] {
			ds.parent[root1] = root2
		} else {
			ds.parent[root2] = root1
			ds.rank[root1]++
		}
	}
}

// kruskal 算法实现
func kruskal(graph *Graph) map[*Edge]struct{} {
	var nodes []*Node
	for _, node := range graph.Nodes {
		nodes = append(nodes, node)
	}
	ds := NewDisjointSet(nodes)
	priorityQueue := NewPriorityQueue()
	for edge, _ := range graph.Edges {
		heap.Push(priorityQueue, edge)
	}
	result := make(map[*Edge]struct{})
	for !priorityQueue.IsEmpty() {
		edge := heap.Pop(priorityQueue).(*Edge)
		if ds.Find(edge.From) != ds.Find(edge.To) {
			result[edge] = struct{}{}
			ds.Union(edge.From, edge.To)
		}
	}

	return result
}
