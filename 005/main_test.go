package main

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	fmt.Println("===========Case 1============")
	whole := "a"
	ret := LongestPalindrome(whole)
	fmt.Println(ret)
	if ret != "a" {
		t.Errorf("expect to receive a but received %s", ret)
	}

	fmt.Println("===========Case 2============")
	whole = "aba"
	ret = LongestPalindrome(whole)
	fmt.Println(ret)
	if ret != "aba" {
		t.Errorf("expect to receive a but received %s", ret)
	}

	fmt.Println("===========Case 3============")
	whole = "abac"
	ret = LongestPalindrome(whole)
	fmt.Println(ret)
	if ret != "aba" {
		t.Errorf("expect to receive a but received %s", ret)
	}

	fmt.Println("===========Case 4============")
	whole = "cbbd"
	ret = LongestPalindrome(whole)
	fmt.Println(ret)
	if ret != "bb" {
		t.Errorf("expect to receive bb but received %s", ret)
	}

	fmt.Println("===========Case 5============")
	whole = "babad"
	ret = LongestPalindrome(whole)
	fmt.Println(ret)
	if ret != "bab" {
		t.Errorf("expect to receive bab but received %s", ret)
	}

	fmt.Println("===========Case 6============")
	whole = "bb"
	ret = LongestPalindrome(whole)
	fmt.Println(ret)
	if ret != "bb" {
		t.Errorf("expect to receive bb but received %s", ret)
	}
	fmt.Println("===========Case 7============")
	whole = "ac"
	ret = LongestPalindrome(whole)
	fmt.Println(ret)
	if ret != "a" && ret != "c" {
		t.Errorf("expect to receive a or c but received %s", ret)
	}
	fmt.Println("===========Case 8============")
	whole = "cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
	ret = LongestPalindrome(whole)

	if len(whole) != len(ret) {
		t.Errorf("%d not equal %d", len(whole), len(ret))
	}
	fmt.Println(ret)

}
