package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	A   []int
	B   []int
	ret []int
}

var testCases = []testCase{
	{
		A:   []int{1, 2},
		B:   []int{2, 3},
		ret: []int{1, 2},
	},
	{
		A:   []int{1, 2, 5},
		B:   []int{2, 4},
		ret: []int{5, 4},
	},
	{
		A:   []int{2},
		B:   []int{1, 3},
		ret: []int{2, 3},
	},
	{
		A:   []int{1, 1},
		B:   []int{2, 2},
		ret: []int{1, 2},
	},
}

func TestFairCandySwap(t *testing.T) {
	for _, v := range testCases {
		ret := FairCandySwap(v.A, v.B)
		fmt.Println(ret)
	}
}
