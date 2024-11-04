package main

import "fmt"

/**
*
* Author: Marek
* Date: 2023-09-22 11:52
* Email: 364021318@qq.com
*
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) AddNode(data int) {
	if l.Val == 0 {
		l.Val = data
	} else {
		if l.Next == nil {
			l.Next = &ListNode{Val:data}
			return
		}
		current := l.Next
		for current != nil {
			if current.Next == nil { // 当最后一个值时，就把新的值写入到他的next的位置
				current.Next = &ListNode{Val:data}
				break
			}
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2 // 在两条链表上的指针
	// 虚拟头结点（构建新链表时的常用技巧）
	dummy := &ListNode{-1, nil}
	// 指针 p 负责构建新链表
	p := dummy
	// 记录进位
	carry := 0
	// 开始执行加法，两条链表走完且没有进位时才能结束循环
	for p1 != nil || p2 != nil || carry > 0 {
		// 先加上上次的进位
		val := carry
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		// 处理进位情况
		carry = val / 10
		val = val % 10
		// 构建新节点
		p.Next = &ListNode{val, nil}
		p = p.Next
	}
	// 返回结果链表的头结点(去除虚拟头结点)
	return dummy.Next
}

func main() {
	l1 := new(ListNode)
	l1.AddNode(2)
	l1.AddNode(4)
	l1.AddNode(3)
	l2 := new(ListNode)
	l2.AddNode(5)
	l2.AddNode(6)
	l2.AddNode(4)
	res := addTwoNumbers(l1, l2)
	// 遍历链表并执行输出
	current := res
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}