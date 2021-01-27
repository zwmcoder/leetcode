package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   int
	ret string
}

var testCases = []testCase{
	{
		s:   3,
		ret: "III",
	},
	{
		s:   4,
		ret: "IV",
	},
	{
		s:   9,
		ret: "IX",
	},
	{
		s:   58,
		ret: "LVIII",
	},
	{
		s:   1994,
		ret: "MCMXCIV",
	},
}

func TestToRoman(t *testing.T) {
	for _, v := range testCases {
		fmt.Println(v)
		ret := IntToRoman(v.s)
		if ret != v.ret {
			t.Errorf("Error! expect %v but received %v", v.ret, ret)
		}
	}
}
