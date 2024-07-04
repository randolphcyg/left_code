package main

import (
	"errors"
	"fmt"
)

/*自定义栈*/

type Stack struct {
	elements []*Node
}

func NewStack() *Stack {
	return &Stack{
		elements: []*Node{},
	}
}

// Push 将节点压入栈中
func (s *Stack) Push(node *Node) {
	s.elements = append(s.elements, node)
}

// Pop 从栈中弹出节点
func (s *Stack) Pop() (*Node, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	// 获取栈顶节点
	top := s.elements[len(s.elements)-1]
	// 移除栈顶节点
	s.elements = s.elements[:len(s.elements)-1]
	return top, nil
}

// Peek 获取栈顶节点而不移除它
func (s *Stack) Peek() (*Node, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size 获取栈的大小
func (s *Stack) Size() int {
	return len(s.elements)
}

// 图的深度优先遍历
func dfs(node *Node) {
	if node == nil {
		return
	}
	stack := NewStack()                 // 创建一个栈
	visited := make(map[*Node]struct{}) // 用于标记已访问过的节点
	stack.Push(node)
	visited[node] = struct{}{}
	fmt.Println(node.Val)
	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		for _, next := range cur.Nexts {
			if _, ok := visited[next]; !ok {
				stack.Push(cur)
				stack.Push(next)
				visited[next] = struct{}{}
				fmt.Println(next.Val)
				break
			}
		}
	}
}
