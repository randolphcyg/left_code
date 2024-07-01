package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
	Rand *ListNode
}

// println 打印单链表
func (head *ListNode) println() {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Rand != nil {
			fmt.Print("(", cur.Rand.Val, ")")
		}
		if cur.Next != nil {
			fmt.Print("->")
		}
		cur = cur.Next
	}
	fmt.Print("->nil")
	fmt.Println()
}

// 用哈希表
func copyListWithRand1(head *ListNode) *ListNode {
	listMap := make(map[*ListNode]*ListNode)
	cur := head
	for cur != nil {
		listMap[cur] = &ListNode{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		// cur 老
		// listMap[cur] 新
		listMap[cur].Next = listMap[cur.Next]
		listMap[cur].Rand = listMap[cur.Rand]
		cur = cur.Next
	}
	return listMap[head]
}

func copyListWithRand2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head

	// 复制每个节点并插入到原节点之后
	for cur != nil {
		next := cur.Next
		cur.Next = &ListNode{Val: cur.Val, Next: next}
		cur = next
	}

	// 设置每个复制节点的 Rand 指针
	cur = head
	for cur != nil {
		// 只关心克隆节点的rand指针如何设置
		if cur.Rand != nil {
			cur.Next.Rand = cur.Rand.Next
		}
		cur = cur.Next.Next
	}

	// 拆分原链表和复制链表
	cur = head
	copyHead := head.Next
	for cur != nil {
		copyNode := cur.Next
		cur.Next = copyNode.Next
		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}
		cur = cur.Next
	}
	return copyHead
}

func main() {
	/*方法1*/

	// 创建链表 1->2->3，并设置 Rand 指针
	head1 := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	third := &ListNode{Val: 3}

	head1.Next = second
	second.Next = third

	head1.Rand = third  // 1 的 Rand 指向 3
	second.Rand = head1 // 2 的 Rand 指向 1

	fmt.Print("head1 original list: ")
	head1.println()

	// 复制链表
	copiedHead := copyListWithRand1(head1)
	fmt.Print("head1 copied list: ")
	copiedHead.println()

	/*方法2*/

	// 创建链表 1->2->3，并设置 Rand 指针
	head2 := &ListNode{Val: 1}
	second2 := &ListNode{Val: 2}
	third2 := &ListNode{Val: 3}

	head2.Next = second2
	second2.Next = third2

	head2.Rand = third2  // 1 的 Rand 指向 3
	second2.Rand = head2 // 2 的 Rand 指向 1

	fmt.Print("head2 original list: ")
	head2.println()

	// 复制链表
	copiedHead2 := copyListWithRand2(head2)
	fmt.Print("copiedHead2 copied list: ")
	copiedHead2.println()
}
