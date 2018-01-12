package numbers

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 存储
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := new(ListNode)
	temp, old := r, 0

	for {

		l1 = nonil(l1)
		l2 = nonil(l2)

		n := l1.Val + l2.Val + old
		if n > 9 {
			n -= 10
			old = 1
		} else {
			old = 0
		}

		temp.Val = n

		// 准备下一个
		l1 = next(l1)
		l2 = next(l2)

		if l1 == nil && l2 == nil {
			break
		}

		// 下一节点
		temp.Next = new(ListNode)
		temp = temp.Next
	}

	if old == 1 {
		temp.Next = &ListNode{Val: old}
	}

	return r
}

// 下一个节点
func next(ln *ListNode) *ListNode {
	if ln != nil {
		return ln.Next
	}
	return nil
}

// 节点不能为nil
func nonil(ln *ListNode) *ListNode {
	if ln != nil {
		return ln
	}
	return &ListNode{Val: 0}
}

// 生成ListNode
func makeListNode(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	ln := new(ListNode)
	ln.Val = s[0]

	temp := ln
	for i := 1; i < len(s); i++ {
		temp.Next = &ListNode{Val: s[i]}
		temp = temp.Next
	}

	return ln
}

// 打印ListNode
func printLN(ln *ListNode) {
	if ln == nil {
		return
	}
	i := 1
	for {
		fmt.Printf("%d:%d\n", i, ln.Val)
		fmt.Printf("%d:%+v\n", i, ln.Next)
		if ln.Next != nil {
			ln = ln.Next
			i++
		} else {
			break
		}
	}
}
