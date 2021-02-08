package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   int
	ret []string
}

var testCases = []testCase{
	//	{
	//		s:   "()",
	//		ret: true,
	//	},
	{
		s:   3,
		ret: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
	{
		s:   2,
		ret: []string{"(())", "()()"},
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
}

func TestIsValid(t *testing.T) {
	for _, v := range testCases {
		fmt.Println(v)
		ret := GenerateParenthesis(v.s)
		fmt.Println(ret)
	}
}
