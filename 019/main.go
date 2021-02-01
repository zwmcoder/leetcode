package main

//	"fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	rawHead := head

	for ; n > 0 && head.Next != nil; n-- {
		head = head.Next
	}

	// n == 1, means n equal the length of the list, the head node been removed
	if n == 1 {
		return rawHead.Next
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
