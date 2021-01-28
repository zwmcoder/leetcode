package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   string
	ret int
}

var testCases = []testCase{
	{
		s:   "III",
		ret: 3,
	},
	{
		s:   "IV",
		ret: 4,
	},
	{
		s:   "IX",
		ret: 9,
	},
	{
		s:   "LVIII",
		ret: 58,
	},
	{
		s:   "MCMXCIV",
		ret: 1994,
	},
}

func TestIsMatch(t *testing.T) {
	for _, v := range testCases {
		fmt.Println(v)
		ret := RomanToInt(v.s)
		if ret != v.ret {
			t.Errorf("Error! expect %v but received %v", v.ret, ret)
		}
	}
}
