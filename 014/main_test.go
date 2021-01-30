package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   []string
	ret string
}

var testCases = []testCase{
	{
		s:   []string{"flower", "flow", "flight"},
		ret: "fl",
	},
	{
		s:   []string{"dog", "racecar", "car"},
		ret: "",
	},
	{
		s:   []string{"ab", "a"},
		ret: "a",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, v := range testCases {
		fmt.Println(v)
		ret := LongestCommonPrefix(v.s)
		if ret != v.ret {
			t.Errorf("Error! expect %v but received %v", v.ret, ret)
		}
	}
}
