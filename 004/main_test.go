package main

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Println("===========Case 1============")
	a := []int{2}
	b := []int{1, 3}
	ret := FindMedianSortedArrays(a, b)

	if ret != 2 {
		t.Errorf("expect 2 but received %f", ret)
	}

	fmt.Println("===========Case 2============")
	a = []int{}
	b = []int{1}
	ret = FindMedianSortedArrays(a, b)

	if ret != 1 {
		t.Errorf("expect 1 but received %f", ret)
	}

	fmt.Println("===========Case 3============")
	a = []int{1, 2}
	b = []int{3, 4}
	ret = FindMedianSortedArrays(a, b)

	if ret != 2.5 {
		t.Errorf("expect 1 but received %f", ret)
	}

}
