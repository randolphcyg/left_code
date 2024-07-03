package main

import "math"

// dijkstra 算法实现
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
			distanceMap[toNode] = min(distanceMap[toNode], distance+edge.Weight)
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
