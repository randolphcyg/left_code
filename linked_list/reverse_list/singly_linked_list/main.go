package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// println 打印单链表
func (head *ListNode) println() {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print("->")
		}
		cur = cur.Next
	}
	fmt.Print("->nil")
	fmt.Println()
}

// reverseListRecursive 翻转链表 递归
// 时O(n) 空O(n)
func reverseListRecursive(head *ListNode) *ListNode {
	return reverseHelper(nil, head)
}

// 辅助递归函数
func reverseHelper(pre, head *ListNode) *ListNode {
	var next *ListNode
	if head == nil {
		return pre
	}

	next, head.Next = head.Next, pre

	return reverseHelper(head, next)
}

// reverseList 翻转链表 双指针
// 时O(n) 空O(1)
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

func main() {
	/*方法1*/

	// 创建链表 1 -> 2 -> 3 -> 4 -> 5
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Print("head1: ")
	head1.println()

	// 链表反转 双指针
	reversedHead1 := reverseList(head1)
	reversedHead1.println()

	/*方法2*/

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Print("head2: ")
	head2.println()

	// 链表反转 递归
	reversedHead2 := reverseListRecursive(head2)
	reversedHead2.println()
}
