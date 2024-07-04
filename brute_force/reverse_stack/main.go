package main

import "fmt"

/*
给你一个栈，请你逆序这个栈，不能申请额外的数据结构，只能使用递归函数。 如何实现?
*/

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Pop from empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func reverse(stack *Stack) {
	if stack.IsEmpty() {
		return
	}
	i := getAndRemoveLastElement(stack)
	reverse(stack)
	stack.Push(i)
}

func getAndRemoveLastElement(stack *Stack) int {
	result := stack.Pop()
	if stack.IsEmpty() {
		return result
	} else {
		last := getAndRemoveLastElement(stack)
		stack.Push(result)
		return last
	}
}

func main() {
	s1 := &Stack{}
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	s1.Push(4)
	s1.Push(5)
	reverse(s1)
	for !s1.IsEmpty() {
		fmt.Println(s1.Pop())
	}
}
