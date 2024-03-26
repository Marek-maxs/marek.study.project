package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-01-06 21:30
* Email: 364021318@qq.com
*
 */

type ListNode struct {
	Val int
	Next *ListNode
}

// 主函数
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{Val: -1, Next: head}
	// 删除倒数第 n 个， 要先找倒数第 n + 1 个节点
	x := findFromEnd(dummy, n + 1)
	// 删除倒数第 n 个节点
	x.Next = x.Next.Next
	return dummy.Next
}

// 返回链表的倒数第 k 个节点
func findFromEnd(head *ListNode, k int) *ListNode  {
	p1, p2 := head, head
	// p1 先走 k 步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	// P1 和 p2 同时走 n -k 步
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// p2 现在指向第 n - k个节点
	return p2
}

func main()  {
	head := new(ListNode)
	head.Val = 1
	head.Next = &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	res := removeNthFromEnd(head, 2)
	fmt.Println(res)
}