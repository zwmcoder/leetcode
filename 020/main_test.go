package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   string
	ret bool
}

var testCases = []testCase{
	//	{
	//		s:   "()",
	//		ret: true,
	//	},
	{
		s:   "()[]{}",
		ret: true,
	},
	//	{
	//		s:   "(]",
	//		ret: false,
	//	},
	//	{
	//		s:   "(]",
	//		ret: false,
	//	},
	//	{
	//		s:   "([)]",
	//		ret: false,
	//	},
	//	{
	//		s:   "{[]}",
	//		ret: true,
	//	},
	{
		s:   "(){}}{",
		ret: false,
	},
}

func TestIsValid(t *testing.T) {
	for _, v := range testCases {
		fmt.Println(v)
		ret := IsValid(v.s)
		if ret != v.ret {
			t.Errorf("expect receive %v, but received %v", v.ret, ret)

		}
	}
}
