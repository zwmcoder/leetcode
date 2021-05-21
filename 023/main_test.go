package main

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	val1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	val2 := &ListNode{
		Val:  4,
		Next: nil,
	}
	val3 := &ListNode{
		Val:  5,
		Next: nil,
	}

	val11 := &ListNode{
		Val:  1,
		Next: nil,
	}
	val12 := &ListNode{
		Val:  3,
		Next: nil,
	}
	val13 := &ListNode{
		Val:  4,
		Next: nil,
	}

	val21 := &ListNode{
		Val:  2,
		Next: nil,
	}
	val22 := &ListNode{
		Val:  6,
		Next: nil,
	}

	val1.Next = val2
	val2.Next = val3

	val11.Next = val12
	val12.Next = val13

	val21.Next = val22

	//	lists := []*ListNode{val1, val11, val21}
	//	fmt.Println("#################")
	//	ret := MergeKLists(lists)
	//	for ; ret != nil; ret = ret.Next {
	//		fmt.Println(ret.Val)
	//	}

	lists := []*ListNode{}
	fmt.Println("#################")
	ret := MergeKLists(lists)
	fmt.Println(ret)
}
