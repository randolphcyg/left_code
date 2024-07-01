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

// 按照 pivot 分割链表
func partition(head *ListNode, pivot int) *ListNode {
	if head == nil {
		return nil
	}

	var smallerHead, smallerTail *ListNode
	var equalHead, equalTail *ListNode
	var largerHead, largerTail *ListNode

	// 遍历链表，分割成三部分
	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val < pivot {
			if smallerHead == nil {
				smallerHead = head
				smallerTail = head
			} else {
				smallerTail.Next = head
				smallerTail = head
			}
		} else if head.Val == pivot {
			if equalHead == nil {
				equalHead = head
				equalTail = head
			} else {
				equalTail.Next = head
				equalTail = head
			}
		} else {
			if largerHead == nil {
				largerHead = head
				largerTail = head
			} else {
				largerTail.Next = head
				largerTail = head
			}
		}
		head = next
	}

	// 连接三部分
	if smallerTail != nil {
		smallerTail.Next = equalHead
		equalTail = smallerTail
	}

	if equalTail != nil {
		equalTail.Next = largerHead
	}

	if smallerHead != nil {
		return smallerHead
	} else if equalHead != nil {
		return equalHead
	} else {
		return largerHead
	}
}

func main() {
	// 创建链表：1 -> 4 -> 3 -> 2 -> 5 -> 2 -> nil
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}

	fmt.Println("Original list:")
	head.println()

	pivot := 3
	newHead := partition(head, pivot)

	fmt.Println("Partitioned list:")
	newHead.println()
}
