package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, index, tmp *ListNode

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		head = l1
		if l1.Next == nil {
			l1.Next = l2
			return head
		}
		l1 = l1.Next
	} else {
		head = l2
		if l2.Next == nil {
			l2.Next = l1
			return head
		}
		l2 = l2.Next
	}

	index = head
	for l1 != nil || l2 != nil {
		//	fmt.Println(l1.Val, l2.Val)
		if l1.Val <= l2.Val {
			// l1 reach the end, and v1.Val is smaller, so set cursour index.Next to l1
			// And set l1.Next to l2, and can go out of this loop
			if l1.Next == nil {
				index.Next = l1
				l1.Next = l2
				break
			}
			tmp = l1
			l1 = l1.Next

		} else {
			if l2.Next == nil {
				index.Next = l2
				l2.Next = l1
				break
			}
			tmp = l2
			l2 = l2.Next
		}
		index.Next = tmp
		index = index.Next

		fmt.Println("###index ", index.Val)
	}

	return head
}
