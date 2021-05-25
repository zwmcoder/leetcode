package main

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	fmt.Println("===========Case 1============")
	b := []int{1, 2, 3}
	ret := Permute(b)

	fmt.Println(ret)

}
