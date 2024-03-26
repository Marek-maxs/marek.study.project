package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-01-06 22:34
* Email: 364021318@qq.com
*
 */

type ListNode struct {
	Val int
	Next *ListNode
	Number int
}

func (n *ListNode) Insert(value int) {
	newNode := &ListNode{Val:value, Next:nil}

	if n.Next == nil { // 如果当前节点没有后继节点，则将新节点设置为第一
		if n.Number == 0 {
			n.Val = value
			n.Number++
		} else {
			n.Next = newNode
			n.Number++
		}

		return
	}

	current := n.Next
	for ; current != nil; current = current.Next {
		if current.Next == nil { // 如果已经到达最后一个节点，则直接将新
			current.Next = newNode
			break
		}
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode  {
	// 虚拟关结点
	dummy := &ListNode{}
	p := dummy

	for l1 != nil && l2 != nil {
		// 比较 p1 和 p2 两个指针
		// 将值较小的节点接到 p 指针
		if l1.Val <= l2.Val { // 取较小的节点添加到结果链表
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		// p 指针不断前进
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return dummy.Next
}

func removeZeroes(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Val == 0 {
		current = current.Next
	}

	if current == nil { // 如果链表为空或只包含0节点，则直接返回 nil
		return nil
	}

	// 递归删除剩余的0节点
	current.Next = removeZeroes(current.Next)

	return head
}

func main() {
	head1 := &ListNode{}
	head1.Insert(1)
	head1.Insert(2)
	head1.Insert(3)

	head2 := &ListNode{}
	head2.Insert(1)
	head2.Insert(2)
	head2.Insert(3)

	res := mergeTwoLists(head1, head2)

	for res != nil {
		fmt.Println(res.Val)

		res = res.Next
	}
}