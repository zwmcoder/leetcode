package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	val1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	val2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	val3 := &ListNode{
		Val:  4,
		Next: nil,
	}

	val4 := &ListNode{
		Val:  1,
		Next: nil,
	}
	val5 := &ListNode{
		Val:  3,
		Next: nil,
	}
	val6 := &ListNode{
		Val:  7,
		Next: nil,
	}

	val1.Next = val2
	val2.Next = val3
	val4.Next = val5
	val5.Next = val6

	//	fmt.Println("#################")
	//	ret := MergeTwoLists(val1, val4)
	//	for ; ret != nil; ret = ret.Next {
	//		fmt.Println(ret.Val)
	//	}
	fmt.Println("#################")
	ret := MergeTwoLists(val3, val6)
	for ; ret != nil; ret = ret.Next {
		fmt.Println(ret.Val)
	}
}
