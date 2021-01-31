package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   string
	ret []string
}

var testCases = []testCase{
	{
		s:   "23",
		ret: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
	{
		s:   "",
		ret: []string{},
	},
	{
		s:   "2",
		ret: []string{"a", "b", "c"},
	},
}

func TestThreeSumClosest(t *testing.T) {
	for _, v := range testCases {
		ret := LetterCombinations(v.s)
		fmt.Println(ret)
	}
}
