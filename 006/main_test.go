package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	fmt.Println("===========Case 1============")

	s := "PAYPALISHIRING"
	ret := Convert(s, 4)
	fmt.Println(ret)
	if ret != "PINALSIGYAHRPI" {
		t.Errorf("Expect receive \"PINALSIGYAHRPI\", but received %s", ret)
	}

	ret = Convert(s, 3)
	fmt.Println(ret)
	if ret != "PAHNAPLSIIGYIR" {
		t.Errorf("Expect receive \"PAHNAPLSIIGYIR\", but received %s", ret)
	}
}
