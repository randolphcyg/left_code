package main

import "fmt"

// 图的广度(宽度)优先遍历
func bfs(node *Node) {
	if node == nil {
		return
	}
	queue := make([]*Node, 0)           // 创建一个队列
	visited := make(map[*Node]struct{}) // 用于标记已访问过的节点
	queue = append(queue, node)
	visited[node] = struct{}{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur.Val)
		for _, next := range cur.Nexts {
			if _, ok := visited[next]; !ok {
				visited[next] = struct{}{}
				queue = append(queue, next)
			}
		}
	}
}
