package main

import "fmt"

/*
题目描述
定义构造三叉搜索树规则如下：

每个节点都存有一个数，当插入一个新的数时，从根节点向下寻找，直到找到一个合适的空节点插入。查找的规则是：

如果数小于节点的数减去500，则将数插入节点的左子树

如果数大于节点的数加上500，则将数插入节点的右子树

否则，将数插入节点的中子树

给你一系列数，请按以上规则，按顺序将数插入树中，构建出一棵三叉搜索树，最后输出树的高度。

输入描述
第一行为一个数 N，表示有 N 个数，1 ≤ N ≤ 10000

从第二行开始输入N行，每行一个整数，每个数的范围为[1,10000]

输出描述
输出树的高度（根节点的高度为1）
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Mid   *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Add(val int) *TreeNode {
	node := &TreeNode{Val: val}

	if t == nil {
		return node
	}

	cur := t
	for {
		if val < cur.Val-500 {
			if cur.Left == nil {
				cur.Left = node
				break
			} else {
				cur = cur.Left
			}
		} else if val > cur.Val+500 {
			if cur.Right == nil {
				cur.Right = node
				break
			} else {
				cur = cur.Right
			}
		} else {
			if cur.Mid == nil {
				cur.Mid = node
				break
			} else {
				cur = cur.Mid
			}
		}
	}
	return t
}

// Height DFS算法求高度
func (t *TreeNode) Height() int {
	if t == nil {
		return 0
	}
	leftHeight := t.Left.Height()
	midHeight := t.Mid.Height()
	rightHeight := t.Right.Height()
	return max(leftHeight, midHeight, rightHeight) + 1
}

/*
输入
9
5000 2000 5000 8000 1800 7500 4500 1400 8100
输出
4
*/
func main() {
	var n int
	fmt.Scan(&n)

	var head int
	fmt.Scan(&head)
	tree := &TreeNode{Val: head}

	for i := 0; i < n-1; i++ {
		var x int
		fmt.Scan(&x)
		tree.Add(x)
	}

	fmt.Println(tree.Height())
}
