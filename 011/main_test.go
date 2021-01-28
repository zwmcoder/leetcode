package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   []int
	ret int
	
}

var testCases = []testCase{
	{
		s:   []int{1,1},
		ret: 1,
	},
	{
		s:   []int{4,3,2,1,4},
		ret: 16,
	},
	{
		s:   []int{1,2,1},
		ret: 2,
	},
	{
		s:   []int{4,3,2,1,4},
		ret: 16,
	},
	{
		s:   []int{1,8,6,2,5,4,8,3,7},
		ret: 49,
	},
}

func TestMaxArea(t *testing.T) {
	for _, v := range testCases {
		fmt.Println(v)
		ret := MaxAreaOn(v.s)
		if ret != v.ret {
			t.Errorf("Error! expect %v but received %v", v.ret, ret)
		}
	}
}
