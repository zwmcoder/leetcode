package main

import (
	"testing"
	"fmt"
)

var testCase = map[string]int{
	"42": 42,
	"-42": -42,
	"   -42": -42,
	"4193 with words": 4193,
	"words and 987": 0,
	"-91283472332": -2147483648,
	"+-12": 0,   //case 1
	"91283472332": 2147483647,
	"00000-42a1234": 0, // case 2
	"3.1415": 3, // case 3
	"-00012a123": -12,  
	"9223372036854775808": 2147483647,
	"-5-": -5,
}
func TestMyAtoi(t *testing.T) {
	for k, v := range testCase {
			fmt.Println(k)
			ret := MyAtoi(k)
			if ret != v {
				t.Errorf("Error! expect %d but received %d", v, ret)
			}
	}
}
