package main

import (
	"fmt"

	"left_code/b_tree/helper"
)

// printTreeByLayer 按层打印二叉树
func printTreeByLayer(root *helper.TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*helper.TreeNode{root}
	for len(queue) > 0 {
		var level []int
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

func main() {
	// 创建一个二叉树
	root := &helper.TreeNode{Val: 1}
	root.Left = &helper.TreeNode{Val: 2}
	root.Right = &helper.TreeNode{Val: 3}
	root.Left.Left = &helper.TreeNode{Val: 4}
	root.Left.Right = &helper.TreeNode{Val: 5}
	root.Right.Left = &helper.TreeNode{Val: 6}
	root.Right.Right = &helper.TreeNode{Val: 7}

	// 使用 printTreeByLayer 方法打印二叉树
	result := printTreeByLayer(root)
	helper.PrintTree(root)

	// 打印结果
	for _, level := range result {
		fmt.Println(level)
	}
}
