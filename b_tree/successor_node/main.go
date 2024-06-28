package main

import (
	"fmt"
	"math"
	"strings"
)

// 在二叉树中找到一个结点的后继结点
/*
该结构比普通二叉树多了一个指向父结点的Parent指针。
只给一个在二叉树中某个结点，请实现返回Node的后继结点的函数
在二叉树的中序遍历的序列中，Node的下一个结点叫做Node的后继结点
*/

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func getSuccessorNode(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.Right != nil {
		return getLeftMost(node.Right)
	} else { // 无右子树
		parent := node.Parent
		for parent != nil && parent.Left != node { // 当前结点是其父结点的右孩子
			node = parent
			parent = node.Parent
		}
		return parent
	}
}

// 从node出发一直往左上
func getLeftMost(node *Node) *Node {
	if node == nil {
		return node
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// PrintTree 打印树形结构
func PrintTree(root *Node) {
	if root == nil {
		return
	}

	// 获取树的深度
	depth := getDepth(root)
	// 计算树的宽度
	width := int(math.Pow(2, float64(depth))) - 1

	// 使用二维切片存储树的节点
	lines := make([][]string, depth*2-1)
	for i := range lines {
		lines[i] = make([]string, width)
		for j := range lines[i] {
			lines[i][j] = " "
		}
	}

	fillLines(root, lines, 0, 0, width-1)

	// 打印每一行
	for _, line := range lines {
		fmt.Println(strings.Join(line, ""))
	}
}

// 递归填充树的节点到二维切片中
func fillLines(node *Node, lines [][]string, depth, left, right int) {
	if node == nil {
		return
	}

	mid := (left + right) / 2
	lines[depth*2][mid] = fmt.Sprintf("%d", node.Val)

	if node.Left != nil {
		lines[depth*2+1][mid-1] = "/"
		fillLines(node.Left, lines, depth+1, left, mid-1)
	}
	if node.Right != nil {
		lines[depth*2+1][mid+1] = "\\"
		fillLines(node.Right, lines, depth+1, mid+1, right)
	}
}

// 计算树的深度
func getDepth(node *Node) int {
	if node == nil {
		return 0
	}
	leftDepth := getDepth(node.Left)
	rightDepth := getDepth(node.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	// 创建一个二叉树
	root := &Node{Val: 6}
	root.Left = &Node{Val: 2, Parent: root}
	root.Right = &Node{Val: 8, Parent: root}
	root.Left.Left = &Node{Val: 0, Parent: root.Left}
	root.Left.Right = &Node{Val: 4, Parent: root.Left}
	root.Left.Right.Left = &Node{Val: 3, Parent: root.Left.Right}
	root.Left.Right.Right = &Node{Val: 5, Parent: root.Left.Right}
	root.Right.Left = &Node{Val: 7, Parent: root.Right}
	root.Right.Right = &Node{Val: 9, Parent: root.Right}

	// 打印树形结构
	PrintTree(root)

	res := getSuccessorNode(root.Left.Right.Left)
	fmt.Println(root.Left.Right.Left.Val, " >> ", res.Val)
}
