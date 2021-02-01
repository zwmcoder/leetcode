package main

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	val1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	val2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	val3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	val4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	val5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	val1.Next = val2
	val2.Next = val3
	val3.Next = val4
	val4.Next = val5

	ret := RemoveNthFromEnd(val1, 2)
	for ; ret != nil; ret = ret.Next {
		fmt.Println(ret.Val)
	}

	val2.Next = nil
	ret = RemoveNthFromEnd(val1, 2)
	for ; ret != nil; ret = ret.Next {
		fmt.Println(ret.Val)
	}

}
