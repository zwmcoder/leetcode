package main

import (
	"testing"
)

type testCase struct {
	s   []int
	t   int
	ret int
}

var testCases = []testCase{
	{
		s:   []int{-1, 2, 1, -4},
		t:   1,
		ret: 2,
	},
	{
		s:   []int{1, 1, 1, 0},
		t:   -100,
		ret: 2,
	},
}

func TestThreeSumClosest(t *testing.T) {
	for _, v := range testCases {
		ret := ThreeSumClosest(v.s, v.t)
		if ret != v.ret {
			t.Errorf("Expect receive %d, but received %d", v.ret, ret)
		}
	}
}
