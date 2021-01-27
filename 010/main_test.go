package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s   string
	p   string
	ret bool
}

var testCases = []testCase{
	{
		s:   "aa",
		p:   "a",
		ret: false,
	},
	{
		s:   "aa",
		p:   "a*",
		ret: true,
	},
	{
		s:   "ab",
		p:   ".*",
		ret: true,
	},
	{
		s:   "aab",
		p:   "c*a*b",
		ret: true,
	},
	{
		s:   "mississippi",
		p:   "mis*is*p*.",
		ret: false,
	},
	{
		s:   "aasdfasdfasdfasdfas",
		p:   "aasdf.*asdf.*asdf.*asdf.*s",
		ret: true,
	},
}

func TestIsMatch(t *testing.T) {
	for _, v := range testCases {
		fmt.Println(v)
		ret := IsMatch(v.s, v.p)
		if ret != v.ret {
			t.Errorf("Error! expect %v but received %v", v.ret, ret)
		}
	}
}
