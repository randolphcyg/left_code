package main

// NodeRecord represents a node with its distance from the source
type NodeRecord struct {
	Node     *Node
	Distance int
}

// NewNodeRecord creates a new NodeRecord
func NewNodeRecord(node *Node, distance int) *NodeRecord {
	return &NodeRecord{Node: node, Distance: distance}
}

// NodeHeap represents a priority queue of nodes
type NodeHeap struct {
	Nodes        []*Node
	HeapIndexMap map[*Node]int // 用于维护节点在堆中的索引位置
	DistanceMap  map[*Node]int
	Size         int
}

func NewNodeHeap(size int) *NodeHeap {
	return &NodeHeap{
		Nodes:        make([]*Node, size),
		HeapIndexMap: make(map[*Node]int),
		DistanceMap:  make(map[*Node]int),
		Size:         0,
	}
}

func (nh *NodeHeap) Len() int { return nh.Size }

func (nh *NodeHeap) Less(i, j int) bool {
	return nh.DistanceMap[nh.Nodes[i]] < nh.DistanceMap[nh.Nodes[j]]
}

func (nh *NodeHeap) Swap(i, j int) {
	nh.HeapIndexMap[nh.Nodes[i]] = j
	nh.HeapIndexMap[nh.Nodes[j]] = i
	nh.Nodes[i], nh.Nodes[j] = nh.Nodes[j], nh.Nodes[i]
}

func (nh *NodeHeap) Push(x interface{}) {
	node := x.(*Node)
	nh.HeapIndexMap[node] = nh.Size
	nh.Nodes[nh.Size] = node
	nh.Size++
}

func (nh *NodeHeap) Pop() interface{} {
	nodeRecord := NewNodeRecord(nh.Nodes[0], nh.DistanceMap[nh.Nodes[0]])
	nh.Swap(0, nh.Size-1)
	nh.HeapIndexMap[nh.Nodes[nh.Size-1]] = -1
	delete(nh.DistanceMap, nh.Nodes[nh.Size-1])
	nh.Nodes[nh.Size-1] = nil
	nh.Size--
	nh.heapify(0, nh.Size)
	return nodeRecord
}

func (nh *NodeHeap) IsEmpty() bool {
	return nh.Size == 0
}

func (nh *NodeHeap) IsEntered(node *Node) bool {
	_, exists := nh.HeapIndexMap[node]
	return exists
}

func (nh *NodeHeap) InHeap(node *Node) bool {
	index, exists := nh.HeapIndexMap[node]
	return exists && index != -1
}

func (nh *NodeHeap) insertHeapify(node *Node, index int) {
	for index > 0 && nh.DistanceMap[nh.Nodes[index]] < nh.DistanceMap[nh.Nodes[(index-1)/2]] {
		nh.Swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (nh *NodeHeap) heapify(index, size int) {
	left := index*2 + 1
	for left < size {
		smallest := left
		if left+1 < size && nh.DistanceMap[nh.Nodes[left+1]] < nh.DistanceMap[nh.Nodes[left]] {
			smallest = left + 1
		}
		if nh.DistanceMap[nh.Nodes[smallest]] >= nh.DistanceMap[nh.Nodes[index]] {
			break
		}
		nh.Swap(index, smallest)
		index = smallest
		left = index*2 + 1
	}
}

func (nh *NodeHeap) AddOrUpdateOrIgnore(node *Node, distance int) {
	if nh.InHeap(node) {
		nh.DistanceMap[node] = min(nh.DistanceMap[node], distance)
		nh.insertHeapify(node, nh.HeapIndexMap[node])
	} else if !nh.IsEntered(node) {
		if nh.Size < len(nh.Nodes) {
			nh.Nodes[nh.Size] = node
		} else {
			nh.Nodes = append(nh.Nodes, node)
		}
		nh.DistanceMap[node] = distance
		nh.HeapIndexMap[node] = nh.Size
		nh.insertHeapify(node, nh.Size)
		nh.Size++
	}
}

// dijkstra2 算法实现 用小根堆改进后
// 从head出发，所有能到达的节点，生成到达每个节点的最小路径记录并返回
func dijkstra2(head *Node, size int) map[*Node]int {
	nodeHeap := NewNodeHeap(size)
	nodeHeap.AddOrUpdateOrIgnore(head, 0)
	result := make(map[*Node]int)
	for !nodeHeap.IsEmpty() {
		record := nodeHeap.Pop().(*NodeRecord)
		cur := record.Node
		distance := record.Distance
		for _, edge := range cur.Edges {
			nodeHeap.AddOrUpdateOrIgnore(edge.To, edge.Weight+distance)
		}
		result[cur] = distance
	}
	return result
}
