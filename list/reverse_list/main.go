package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 翻转链表-双指针
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

// reverseListHelp 翻转链表-递归
func reverseListHelp(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	var next *ListNode
	if head == nil {
		return pre
	}
	next, head.Next = head.Next, pre

	return help(head, next)
}

func main() {
	// 创建一个单链表
	n1 := new(ListNode)
	n1.Val = 1

	n2 := new(ListNode)
	n2.Val = 2

	n3 := new(ListNode)
	n3.Val = 3

	n4 := new(ListNode)
	n4.Val = 4

	n5 := new(ListNode)
	n5.Val = 5

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	fmt.Println(n1)

	// 链表反转-双指针
	reverseN1 := reverseList(n1)
	fmt.Println(reverseN1)

	// 链表反转-递归
	reverseListHelpN1 := reverseListHelp(n1)
	fmt.Println(reverseListHelpN1)

	fmt.Print((1 + 2) / 2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 尾节点
	var head, tail *ListNode
	carry := 0
	// 遍历俩单链表
	for l1 != nil || l2 != nil {
		// 初始化临时变量 存储俩链表当前计算节点的值
		n1, n2 := 0, 0

		// 如果链表当前计算节点不为空 则取到节点的值
		if l1 != nil {
			n1 = l1.Val
			// 节点后移
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		//  将节点值相加 并加上进位值 并重置下一个节点计算的进位值
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		// 头结点不变的情况下指针可以右移，所以说tail相当于临时节点，理解成尾节点也可以
		// 因为此时新链表中只有一个节点，所以头结点和尾结点都指向同一个元素
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 第二个节点后开始逐渐往尾结点增加元素
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}

	}

	// 最后一位的余数加到链表最后
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}
