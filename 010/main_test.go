package main

import (
	"fmt"
	"testing"
)

var testCase = map[int]bool{
	121:  true,
	-121: false,
	10:   false,
}

func TestIsPalindrome(t *testing.T) {
	for k, v := range testCase {
		fmt.Println(k)
		ret := IsPalindrome(k)
		if ret != v {
			t.Errorf("Error! expect %v but received %v", v, ret)
		}
	}
}
