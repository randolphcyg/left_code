package main

import "fmt"

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

// 反转链表并返回新的头节点
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// 判断链表是否为回文结构（方法1：使用栈）
func isPalindromeMethod1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 使用快慢指针找到链表的中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 将链表的右半部分节点压入栈中
	var stack []*ListNode
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}

	// 从头开始遍历，比较栈中的节点
	for head != nil && len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if head.Val != top.Val {
			return false
		}
		head = head.Next
	}

	return true
}

// 判断链表是否为回文结构（方法2：改指针）
/*
使用快慢指针找到链表的中间节点。
反转链表的后半部分。
比较链表的前半部分和反转后的后半部分是否相等。
恢复链表的后半部分（以保持链表的原始结构）。
返回比较结果。
*/
func isPalindromeMethod2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 使用快慢指针找到链表的中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转链表的后半部分
	secondHalfStart := reverseList(slow)
	copySecondHalfStart := secondHalfStart // 保存反转后链表头，用于恢复链表

	// 比较前半部分和反转后的后半部分
	p1, p2 := head, secondHalfStart
	result := true
	for p2 != nil {
		if p1.Val != p2.Val {
			result = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 恢复链表的后半部分
	reverseList(copySecondHalfStart)

	return result
}

func main() {
	/*方法1*/

	// 创建链表：1 -> 2 -> 3 -> 2 -> 1 -> nil
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 1}

	fmt.Println("head:")
	head.println()

	if isPalindromeMethod1(head) {
		fmt.Println("head: The list is a palindrome.")
	} else {
		fmt.Println("head: The list is not a palindrome.")
	}

	/*方法2*/

	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 3}
	head2.Next.Next.Next = &ListNode{Val: 2}
	head2.Next.Next.Next.Next = &ListNode{Val: 1}

	fmt.Println("head2:")
	head2.println()

	if isPalindromeMethod2(head2) {
		fmt.Println("head2: The list is a palindrome.")
	} else {
		fmt.Println("head2: The list is not a palindrome.")
	}
}
