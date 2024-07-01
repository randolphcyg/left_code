package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印两个有序链表的公共部分
// 双指针的方法，谁小谁移动，相等则打印并同时移动两个指针，有一个指针越界了就停止
func printCommonParts(head1, head2 *ListNode) {
	p1, p2 := head1, head2

	for p1 != nil && p2 != nil {
		if p1.Val == p2.Val {
			fmt.Print(p1.Val, " ")
			p1 = p1.Next
			p2 = p2.Next
		} else if p1.Val < p2.Val {
			p1 = p1.Next
		} else {
			p2 = p2.Next
		}
	}
	fmt.Println()
}

func main() {
	// 创建第一个有序链表：1 -> 2 -> 3 -> 4 -> 5 -> nil
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = &ListNode{Val: 5}

	// 创建第二个有序链表：3 -> 4 -> 5 -> 6 -> 7 -> nil
	head2 := &ListNode{Val: 3}
	head2.Next = &ListNode{Val: 4}
	head2.Next.Next = &ListNode{Val: 5}
	head2.Next.Next.Next = &ListNode{Val: 6}
	head2.Next.Next.Next.Next = &ListNode{Val: 7}

	fmt.Println("Common parts of the two lists are:")
	printCommonParts(head1, head2)
}
