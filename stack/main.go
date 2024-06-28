package main

import (
	"fmt"
)

type Stack []any

func (s *Stack) Push(v any) {
	*s = append(*s, v)
}

func (s *Stack) Pop() any {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Top() any {
	return (*s)[len(*s)-1]
}

func main() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}
