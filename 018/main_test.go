package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s []int
	t int
	//	ret [][]int
}

var testCases = []testCase{
	{
		s: []int{1, 0, -1, 0, -2, 2},
		t: 0,
	},
	{
		s: []int{},
		t: 0,
	},
}

func TestFourSum(t *testing.T) {
	for _, v := range testCases {
		ret := FourSum(v.s, v.t)
		fmt.Println(ret)
	}
}
