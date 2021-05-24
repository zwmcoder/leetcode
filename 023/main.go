package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	min_idx := 0

	var tmp *ListNode

	nodeList := make([]*ListNode, 0)

	if len(lists) == 0 {
		return nil
	}

Loop:
	for {
		var min_val *int
		for _, v1 := range lists {
			if v1 == nil {
				continue
			}
			min_val = &v1.Val

		}
		if min_val == nil {
			break Loop
		}
		for k, v := range lists {
			if v == nil {
				continue
			}
			if v.Val <= *min_val {
				min_val = &v.Val
				min_idx = k
			}
		}
		tmp = lists[min_idx]

		if tmp != nil {
			lists[min_idx] = tmp.Next
		}
		nodeList = append(nodeList, tmp)
	}

	for idx := 0; idx < len(nodeList)-1; idx++ {
		nodeList[idx].Next = nodeList[idx+1]
	}
	if len(nodeList) == 0 {
		return nil
	}
	return nodeList[0]

}
