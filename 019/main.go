package main

//	"fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		if n == 1 {
			return nil
		} else {
			return nil
		}
	}

	rawHead := head
	// n need move one more step, since we need the node 3 not node 4
	// for the example picture at description
	// So set n >= 0, not n > 0
	for ; n >= 0 && head.Next != nil; n-- {
		head = head.Next
	}

	cursor := rawHead
	for ; head != nil && head.Next != nil; head = head.Next {
		cursor = cursor.Next
	}
	p := cursor.Next
	cursor.Next = cursor.Next.Next
	p.Next = nil

	if cursor == rawHead {
		return cursor
	} else {
		return rawHead
	}
}
