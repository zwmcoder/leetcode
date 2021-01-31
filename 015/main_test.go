package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   []int
	ret [][]int
}

var testCases = []testCase{
	{
		s:   []int{-1, 0, 1, 2, -1, -4},
		ret: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		s:   []int{0, 0, 0},
		ret: [][]int{{0, 0, 0}},
	},
}

func TestThreeSum(t *testing.T) {
	for _, v := range testCases {
		ret := ThreeSum(v.s)
		fmt.Println("result: ", ret)
	}
}
