package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	first := true
	n := len(lists)
	nil_count := 0
	min_idx := 0
	min_val := lists[0].Val

	var head, tmp *ListNode
	for {
		for k, v := range lists {
			if v == nil {
				nil_count++
				if nil_count == n {
					break
				}
				continue
			}
			if v.Val <= min_val {
				min_val = v.Val
				min_idx = k
			}
		}
		tmp = lists[min_idx]
		if first {
			head = tmp
			first = false
		}

		if tmp != nil {
			lists[min_idx] = tmp.Next
		}

	}

	return head

}
