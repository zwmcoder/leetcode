package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println("===========Case 1============")
	s := 123
	ret := Reverse(s)
	fmt.Println(ret)
	if ret != 321 {
		t.Errorf("Expect receive 321, but received %d", ret)
	}

	fmt.Println("===========Case 2============")
	s = 0
	ret = Reverse(s)
	fmt.Println(ret)
	if ret != 0 {
		t.Errorf("Expect receive 0, but received %d", ret)
	}

	fmt.Println("===========Case 3============")
	s = 5
	ret = Reverse(s)
	fmt.Println(ret)
	if ret != 5 {
		t.Errorf("Expect receive 5, but received %d", ret)
	}
	
	fmt.Println("===========Case 3============")
	s = 120
	ret = Reverse(s)
	fmt.Println(ret)
	if ret != 21 {
		t.Errorf("Expect receive 21, but received %d", ret)
	}


}
