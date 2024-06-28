package helper

import (
	"fmt"
	"math"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PrintTree 打印树形结构
func PrintTree(root *TreeNode) {
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
func fillLines(node *TreeNode, lines [][]string, depth, left, right int) {
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
func getDepth(node *TreeNode) int {
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
