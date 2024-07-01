package main

import "fmt"

type ListNode struct {
	Val  int
	Prev *ListNode
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

// 翻转双向链表 双指针
// 时O(n) 空O(1)
func reverseDoublyLinkedList(head *ListNode) *ListNode {
	var newHead *ListNode
	current := head
	for current != nil {
		// 交换 prev 和 next
		current.Prev, current.Next = current.Next, current.Prev
		// 更新 newHead 为当前节点
		newHead = current
		// 移动到下一个节点
		current = current.Prev
	}
	return newHead
}

// 递归反转双向链表的方法
// 时O(n) 空O(n)
func reverseDoublyLinkedListRecursive(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	// 交换 prev 和 next
	node.Prev, node.Next = node.Next, node.Prev

	// 如果交换后的 prev 是 nil，说明此节点是新的头节点
	if node.Prev == nil {
		return node
	}

	// 继续反转下一个节点
	return reverseDoublyLinkedListRecursive(node.Prev)
}

func main() {
	/*方法1*/

	// 创建双向链表 1 <-> 2 <-> 3 <-> 4 <-> 5
	head1 := &ListNode{1, nil, nil}
	second := &ListNode{2, head1, nil}
	head1.Next = second
	third := &ListNode{3, second, nil}
	second.Next = third
	fourth := &ListNode{4, third, nil}
	third.Next = fourth
	fifth := &ListNode{5, fourth, nil}
	fourth.Next = fifth
	fmt.Print("head1: ")
	head1.println()

	// 翻转双向链表 双指针 head1
	reversedHead1 := reverseDoublyLinkedList(head1)
	fmt.Print("Reversed head1: ")
	reversedHead1.println()

	/*方法2*/

	// 创建双向链表 1 <-> 2 <-> 3 <-> 4 <-> 5
	head2 := &ListNode{1, nil, nil}
	second2 := &ListNode{2, head2, nil}
	head2.Next = second2
	third2 := &ListNode{3, second2, nil}
	second2.Next = third2
	fourth2 := &ListNode{4, third2, nil}
	third2.Next = fourth2
	fifth2 := &ListNode{5, fourth2, nil}
	fourth2.Next = fifth2
	fmt.Print("head1: ")
	head2.println()
	// 链表反转 递归 head2
	reversedHead2 := reverseDoublyLinkedListRecursive(head2)
	fmt.Print("Reversed head2: ")
	reversedHead2.println()
}
