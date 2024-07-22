package main

import "fmt"

// DisjointSet 并查集

// V 自定义类型
type V int

// Element 结构体，用于存储值
type Element struct {
	value V
}

// Stack 结构体，用于实现堆栈
type Stack struct {
	elements []*Element
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(ele *Element) {
	s.elements = append(s.elements, ele)
}

// Pop 方法，用于从堆栈弹出元素
func (s *Stack) Pop() *Element {
	if len(s.elements) == 0 {
		var zero *Element
		return zero
	}
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func NewElement(value V) *Element {
	return &Element{value: value}
}

type UnionFindSet struct {
	elementMap map[V]*Element
	// key 某个元素value该元素的父
	fatherMap map[*Element]*Element
	// key 某个集合的代表元素 value该集合的大小
	sizeMap map[*Element]int
}

func NewUnionFindSet(list []V) *UnionFindSet {
	elementMap := make(map[V]*Element)
	fatherMap := make(map[*Element]*Element)
	sizeMap := make(map[*Element]int)
	for _, value := range list {
		ele := NewElement(value)
		elementMap[value] = ele
		fatherMap[ele] = ele
		sizeMap[ele] = 1
	}
	return &UnionFindSet{
		elementMap: elementMap,
		fatherMap:  fatherMap,
		sizeMap:    sizeMap,
	}
}

// 给定一个ele 往上一直找 把代表元素返回
func (s *UnionFindSet) findHead(element *Element) *Element {
	path := NewStack()

	for element != s.fatherMap[element] {
		path.Push(element)
		element = s.fatherMap[element]
	}

	for !path.IsEmpty() {
		s.fatherMap[path.Pop()] = element
	}
	return element
}

func (s *UnionFindSet) isSameSet(a, b V) bool {
	if ele1, ok1 := s.elementMap[a]; ok1 {
		if ele2, ok2 := s.elementMap[b]; ok2 {
			return s.findHead(ele1) == s.findHead(ele2)
		}
	}
	return false
}

func (s *UnionFindSet) union(a, b V) {
	if ele1, ok1 := s.elementMap[a]; ok1 {
		if ele2, ok2 := s.elementMap[b]; ok2 {
			aF := s.findHead(ele1)
			bF := s.findHead(ele2)
			if aF != bF {
				size1, _ := s.sizeMap[aF]
				size2, _ := s.sizeMap[bF]
				big := aF

				if size1 < size2 {
					big = bF
				}

				small := bF
				if big != aF {
					small = aF
				}

				s.fatherMap[small] = big
				s.sizeMap[big] = s.sizeMap[aF] + s.sizeMap[bF]
				delete(s.sizeMap, small)
			}
		}
	}
}

func main() {
	// 创建一个包含元素 1 到 5 的 UnionFindSet
	elements := []V{1, 2, 3, 4, 5}
	ufs := NewUnionFindSet(elements)

	// 合并一些元素
	ufs.union(1, 2)
	ufs.union(2, 3)
	ufs.union(4, 5)

	// 检查一些元素是否属于同一集合
	fmt.Println(ufs.isSameSet(1, 3)) // 输出: true
	fmt.Println(ufs.isSameSet(1, 4)) // 输出: false

	// 合并更多元素
	ufs.union(3, 4)

	// 再次检查元素是否属于同一集合
	fmt.Println(ufs.isSameSet(1, 5)) // 输出: true
	fmt.Println(ufs.isSameSet(2, 5)) // 输出: true
	fmt.Println(ufs.isSameSet(5, 5)) // 输出: true
	fmt.Println(ufs.isSameSet(1, 6)) // 输出: false
}
