package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueNode struct {
	Node  *TreeNode
	Index int
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

// bfs 广度优先遍历
func bfs(root *TreeNode) {
	if root == nil {
		return
	}

	// 创建队列，把根节点放入队列
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// 取出队列的第一个元素
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Val)

		// 把左右子节点放入队列
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

// 宽度优先遍历 求最大宽度
func maxWidthOfBT(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []QueueNode{{Node: root, Index: 0}}
	maxWidth := 1

	for len(queue) > 0 {
		levelLen := len(queue)
		left := queue[0].Index

		for i := 0; i < levelLen; i++ {
			node := queue[0].Node
			index := queue[0].Index
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, QueueNode{Node: node.Left, Index: index * 2})
			}
			if node.Right != nil {
				queue = append(queue, QueueNode{Node: node.Right, Index: index*2 + 1})
			}

			if i == levelLen-1 && index-left+1 > maxWidth {
				maxWidth = index - left + 1
			}
		}
	}

	return maxWidth
}

func isBST(root *TreeNode) bool {
	return isBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isBSTHelper(node *TreeNode, lower int, upper int) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if val <= lower || val >= upper {
		return false
	}
	return isBSTHelper(node.Right, val, upper) && isBSTHelper(node.Left, lower, val)
}

// isCBT 完全二叉树
func isCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 是否遇到过左右两个孩子不双全的结点
	isLeaf := false
	var l, r *TreeNode
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		head := queue[0]
		l, r = head.Left, head.Right
		queue = queue[1:]

		if (isLeaf && (l != nil || r != nil)) || // 若遇到不双全结点，又发现当前结点居然有孩子
			(l == nil && r != nil) { // 条件1 有右无左
			return false
		}

		if l != nil {
			queue = append(queue, l)
		}
		if r != nil {
			queue = append(queue, r)
		} else {
			isLeaf = true
		}
	}
	return true
}

// 是否是平衡二叉树

type returnDataBalancedTree struct {
	IsBalanced bool
	Height     int
}

// isBalancedTree 平衡二叉树
func isBalancedTree(root *TreeNode) bool {
	return processBalancedTree(root).IsBalanced
}

func processBalancedTree(x *TreeNode) *returnDataBalancedTree {
	if x == nil { // base
		return &returnDataBalancedTree{true, 0}
	}

	leftData := processBalancedTree(x.Left)
	rightData := processBalancedTree(x.Right)

	height := max(leftData.Height, rightData.Height) + 1
	isBalanced := leftData.IsBalanced && rightData.IsBalanced && math.Abs(float64(leftData.Height-rightData.Height)) < 2

	return &returnDataBalancedTree{isBalanced, height}
}

type returnDataBST struct {
	IsBst bool
	Min   int
	Max   int
}

func isBST2(root *TreeNode) bool {
	return processBST(root).IsBst
}

func processBST(x *TreeNode) *returnDataBST {
	if x == nil {
		return nil
	}
	leftData := processBST(x.Left)
	rightData := processBST(x.Right)

	minData := x.Val
	maxData := x.Val
	if leftData != nil {
		minData = min(minData, leftData.Min)
		maxData = max(maxData, leftData.Max)
	}
	if rightData != nil {
		minData = min(minData, rightData.Min)
		maxData = max(maxData, rightData.Max)
	}

	isBst := true
	if leftData != nil && (!leftData.IsBst || leftData.Max >= x.Val) {
		isBst = false
	}
	if rightData != nil && (!rightData.IsBst || x.Val >= rightData.Min) {
		isBst = false
	}

	return &returnDataBST{isBst, minData, maxData}
}

// 满二叉树

type returnDataFT struct {
	Height int
	Nodes  int
}

func isFT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	data := processFT(root)

	return data.Nodes == (1<<data.Height - 1)
}

func processFT(x *TreeNode) *returnDataFT {
	if x == nil {
		return &returnDataFT{0, 0}
	}
	leftData := processFT(x.Left)
	rightData := processFT(x.Right)
	height := max(leftData.Height, rightData.Height) + 1
	nodes := leftData.Nodes + rightData.Nodes + 1

	return &returnDataFT{
		Height: height,
		Nodes:  nodes,
	}
}

// o1和o2一定属于head为头的树
// 返回最低公共祖先
/*
另一个可能的优化是利用二叉搜索树的性质（如果有的话）。
如果我们的树是一个二叉搜索树，我们可以利用节点值的大小关系来快速找到最低公共祖先。
从根节点开始，如果o1和o2的值都比当前节点小，那么最低公共祖先一定在左子树，
如果o1和o2的值都比当前节点大，那么最低公共祖先一定在右子树，否则，当前节点就是最低公共祖先。
这种方法的时间复杂度是O(logN)，但它只适用于二叉搜索树。
*/
func lowestCommonAncestor(head, o1, o2 *TreeNode) *TreeNode {
	fatherMap := make(map[*TreeNode]*TreeNode)
	fatherMap[head] = head // 为根节点设置父节点为它自己
	processLCA(head, fatherMap)
	set1 := make(map[*TreeNode]struct{})
	cur := o1
	for cur != fatherMap[cur] {
		set1[cur] = struct{}{}
		cur = fatherMap[cur]
	}
	set1[head] = struct{}{}

	// 这里遍历o2, 判断第一次遇到一个结点在set1中则返回这个节点
	cur = o2
	for _, ok := set1[cur]; !ok; _, ok = set1[cur] {
		cur = fatherMap[cur]
	}
	return cur
}

func processLCA(head *TreeNode, fatherMap map[*TreeNode]*TreeNode) {
	if head == nil {
		return
	}
	if head.Left != nil {
		fatherMap[head.Left] = head
	}
	if head.Right != nil {
		fatherMap[head.Right] = head
	}
	processLCA(head.Left, fatherMap)
	processLCA(head.Right, fatherMap)
}

// 序列化
func serialize(root *TreeNode) string {
	var res []string
	serializeHelper(root, &res)
	return strings.Join(res, ",")
}

func serializeHelper(node *TreeNode, res *[]string) {
	if node == nil {
		*res = append(*res, "null")
		return
	}
	*res = append(*res, strconv.Itoa(node.Val))
	serializeHelper(node.Left, res)
	serializeHelper(node.Right, res)
}

// 反序列化
func deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")
	var idx int
	return deserializeHelper(nodes, &idx)
}

func deserializeHelper(nodes []string, idx *int) *TreeNode {
	if *idx >= len(nodes) || nodes[*idx] == "null" {
		*idx++
		return nil
	}
	val, _ := strconv.Atoi(nodes[*idx])
	node := &TreeNode{Val: val}
	*idx++
	node.Left = deserializeHelper(nodes, idx)
	node.Right = deserializeHelper(nodes, idx)
	return node
}

func main() {
	// 创建一个简单的二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	bfs(root)
	PrintTree(root)
	fmt.Println("max width:", maxWidthOfBT(root)) // 输出：4

	// 搜索二叉树
	// 创建一个搜索二叉树
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}

	// 创建一个不是搜索二叉树的二叉树
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}

	fmt.Println("isBST：", isBST(root1))   // 输出：true
	fmt.Println("isBST2：", isBST2(root1)) // 输出：true
	PrintTree(root1)
	fmt.Println("isBST：", isBST(root2))   // 输出：false
	fmt.Println("isBST2：", isBST2(root2)) // 输出：false
	PrintTree(root2)

	// 完全二叉树
	// 创建一个完全二叉树
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	root3.Right = &TreeNode{Val: 3}
	root3.Left.Left = &TreeNode{Val: 4}
	root3.Left.Right = &TreeNode{Val: 5}
	root3.Right.Left = &TreeNode{Val: 6}

	// 创建一个不是完全二叉树的二叉树
	root4 := &TreeNode{Val: 1}
	root4.Left = &TreeNode{Val: 2}
	root4.Right = &TreeNode{Val: 3}
	root4.Left.Left = &TreeNode{Val: 4}
	root4.Right.Right = &TreeNode{Val: 6}

	fmt.Println("isCBT：", isCBT(root3)) // 输出：true
	PrintTree(root3)
	fmt.Println("isCBT：", isCBT(root4)) // 输出：false
	PrintTree(root4)

	// 创建一个平衡二叉树
	root5 := &TreeNode{Val: 1}
	root5.Left = &TreeNode{Val: 2}
	root5.Right = &TreeNode{Val: 3}
	root5.Left.Left = &TreeNode{Val: 4}
	root5.Left.Right = &TreeNode{Val: 5}
	root5.Right.Left = &TreeNode{Val: 6}

	// 创建一个不是平衡二叉树的二叉树
	root6 := &TreeNode{Val: 1}
	root6.Left = &TreeNode{Val: 2}
	root6.Left.Left = &TreeNode{Val: 3}
	root6.Left.Left.Left = &TreeNode{Val: 4}

	fmt.Println("isBalancedTree: ", isBalancedTree(root5)) // 输出：true
	PrintTree(root5)
	fmt.Println("isBalancedTree: ", isBalancedTree(root6)) // 输出：false
	PrintTree(root6)

	// 创建一个满二叉树
	root7 := &TreeNode{Val: 1}
	root7.Left = &TreeNode{Val: 2}
	root7.Right = &TreeNode{Val: 3}
	root7.Left.Left = &TreeNode{Val: 4}
	root7.Left.Right = &TreeNode{Val: 5}
	root7.Right.Left = &TreeNode{Val: 6}
	root7.Right.Right = &TreeNode{Val: 7}

	// 创建一个不是满二叉树的二叉树
	root8 := &TreeNode{Val: 1}
	root8.Left = &TreeNode{Val: 2}
	root8.Right = &TreeNode{Val: 3}
	root8.Left.Left = &TreeNode{Val: 4}
	root8.Right.Left = &TreeNode{Val: 6}

	fmt.Println("isFT: ", isFT(root7)) // 输出：true
	PrintTree(root7)
	fmt.Println("isFT: ", isFT(root8)) // 输出：false
	PrintTree(root8)

	// 创建一个二叉树
	root9 := &TreeNode{Val: 3}
	root9.Left = &TreeNode{Val: 5}
	root9.Right = &TreeNode{Val: 1}
	root9.Left.Left = &TreeNode{Val: 6}
	root9.Left.Right = &TreeNode{Val: 2}
	root9.Right.Left = &TreeNode{Val: 0}
	root9.Right.Right = &TreeNode{Val: 8}

	node1 := root9.Left  // 5
	node2 := root9.Right // 1

	ancestor := lowestCommonAncestor(root9, node1, node2)
	fmt.Println("ancestor: ", ancestor.Val) // 输出：3
	PrintTree(root9)

	root10 := &TreeNode{Val: 1}
	root10.Left = &TreeNode{Val: 2}
	root10.Right = &TreeNode{Val: 3}
	root10.Right.Left = &TreeNode{Val: 4}
	root10.Right.Right = &TreeNode{Val: 5}

	serialized := serialize(root10)
	fmt.Println("Serialized:", serialized)
	PrintTree(root10)

	deserialized := deserialize(serialized)
	fmt.Println("Deserialized Root Value:", deserialized.Val)
}
