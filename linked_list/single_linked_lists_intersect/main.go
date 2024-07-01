package main

import (
	"fmt"
)

// TODO
type ListNode struct {
	Val  int
	Next *ListNode
}

// println 打印单链表
func (head *ListNode) println() {
	visited := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if visited[cur] {
			fmt.Printf("%d (cycle detected at %p)", cur.Val, cur)
			break
		}
		fmt.Printf("%d(%p)", cur.Val, cur)
		visited[cur] = true
		if cur.Next != nil {
			fmt.Print("->")
		}
		cur = cur.Next
	}
	fmt.Print("->nil")
	fmt.Println()
}

// len 单链表长度
func (head *ListNode) len() int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

// getLoopNode [推荐这种解法]找到两个单链表相交的第一个节点 特性法[分析出特性并利用]
// 俩链表相交，那么相交到结尾部分是共用的
// 长链表先走长度差值步数，然后俩链表一起走，返回第一个相交的结点处
func noLoop2(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}

	// 计算两个链表的长度
	len1 := head1.len()
	len2 := head2.len()

	// 对齐两个链表
	if len1 > len2 {
		for len1 > len2 {
			head1 = head1.Next
			len1--
		}
	} else {
		for len2 > len1 {
			head2 = head2.Next
			len2--
		}
	}

	// 同时遍历两个链表，找到相交的第一个节点
	for head1 != nil && head2 != nil {
		if head1 == head2 {
			return head1
		}
		head1 = head1.Next
		head2 = head2.Next
	}

	return nil
}

// getLoopNode2 双指针法
func noLoop(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}

	p1 := head1
	p2 := head2

	// 当 p1 和 p2 不相等时，继续循环
	for p1 != p2 {
		// p1 走一步，如果走到末尾就转向 head2
		if p1 == nil {
			p1 = head2
		} else {
			p1 = p1.Next
		}

		// p2 走一步，如果走到末尾就转向 head1
		if p2 == nil {
			p2 = head1
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

func getIntersectNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2) // 俩无环链表的相交问题
	}
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, loop1, head2, loop2)
	}
	return nil
}

// getLoopNode 找到链表的入环节点
func getLoopNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	n1 := head.Next      // n1 -> slow
	n2 := head.Next.Next // n1 -> fast
	for n1 != n2 {
		if n2.Next == nil || n2.Next.Next == nil {
			return nil
		}
		n2 = n2.Next.Next
		n1 = n1.Next
	}

	n2 = head // n2 -> walk again from head
	for n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1
}

// bothLoop 俩有环链表相交问题
func bothLoop(head1, loop1, head2, loop2 *ListNode) *ListNode {
	if loop1 == loop2 {
		cur1, cur2 := head1, head2

		// 对齐两个链表在环入口之前的部分
		len1, len2 := 0, 0
		for cur1 != loop1 {
			len1++
			cur1 = cur1.Next
		}
		for cur2 != loop2 {
			len2++
			cur2 = cur2.Next
		}

		cur1, cur2 = head1, head2
		if len1 > len2 {
			for len1 > len2 {
				cur1 = cur1.Next
				len1--
			}
		} else {
			for len2 > len1 {
				cur2 = cur2.Next
				len2--
			}
		}

		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}

		return cur1
	} else {
		cur1 := loop1.Next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1 // 环中相交，返回任意一个环入口
			}
			cur1 = cur1.Next
		}
		return nil // 环中不相交
	}
}

/*
例子1：两个无环链表相交
链表1: 1 -> 2 -> 3 -> 6 -> 7
链表2: 4 -> 5 -> 6 -> 7

例子2：两个无环链表不相交
链表1: 1 -> 2 -> 3
链表2: 4 -> 5 -> 6

例子3：两个有环链表相交
链表1: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 3（回到节点3形成环）
链表2: 7 -> 8 -> 5（与链表1相交） -> 6 -> 3（回到节点3形成环）

例子4：两个有环链表不相交
链表1: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 3（回到节点3形成环）
链表2: 7 -> 8 -> 9 -> 10 -> 8（回到节点8形成环）

例子5：一个有环链表，一个无环链表
链表1: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 3（回到节点3形成环）
链表2: 7 -> 8 -> 9
*/

func main() {
	fmt.Println("## 1.两个无环链表相交")
	// 例子1：两个无环链表相交
	common := &ListNode{Val: 6, Next: &ListNode{Val: 7}}
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: common}}}
	head2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: common}}

	fmt.Print("无环链表1: ")
	head1.println()
	fmt.Print("无环链表2: ")
	head2.println()

	intersection := getIntersectNode(head1, head2)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}

	fmt.Println()
	fmt.Println("## 2.两个无环链表不相交")
	// 例子2：两个无环链表不相交
	head1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	head2 = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}

	fmt.Print("无环链表1: ")
	head1.println()
	fmt.Print("无环链表2: ")
	head2.println()

	intersection = getIntersectNode(head1, head2)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}

	fmt.Println()
	fmt.Println("## 3.两个有环链表相交")
	// 例子3：两个有环链表相交
	loop := &ListNode{Val: 3}
	common = &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: loop}}
	loop.Next = common
	head1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: loop}}
	head2 = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: common}}

	fmt.Print("有环链表1: ")
	head1.println()
	fmt.Print("有环链表2: ")
	head2.println()

	intersection = getIntersectNode(head1, head2)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}

	fmt.Println()
	fmt.Println("## 4.两个有环链表不相交")
	// 例子4：两个有环链表不相交
	loop1 := &ListNode{Val: 3}
	head1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: loop1}}
	loop1.Next = head1.Next

	loop2 := &ListNode{Val: 8}
	head2 = &ListNode{Val: 7, Next: loop2}
	loop2.Next = head2

	fmt.Print("有环链表1: ")
	head1.println()
	fmt.Print("有环链表2: ")
	head2.println()

	intersection = getIntersectNode(head1, head2)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}

	fmt.Println()
	fmt.Println("## 5.一个有环链表，一个无环链表")
	// 例子5：一个有环链表，一个无环链表
	head1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: loop1}}
	head2 = &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}}

	fmt.Print("有环链表: ")
	head1.println()
	fmt.Print("无环链表: ")
	head2.println()

	intersection = getIntersectNode(head1, head2)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}

	fmt.Println()
	fmt.Println("## 6.两个有环链表相交，且入环节点是相同的")
	// 例子：两个有环链表相交，且入环节点是相同的
	loop3 := &ListNode{Val: 3}
	common2 := &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: loop3}}
	loop.Next = common2

	// 构建第一个有环链表
	head1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: loop3}}

	// 构建第二个有环链表
	head2 = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: common}}

	fmt.Print("有环链表1: ")
	head1.println()
	fmt.Print("有环链表2: ")
	head2.println()

	intersection = getIntersectNode(head1, head2)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}
}
