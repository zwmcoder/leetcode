package main

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	ret := CoinChange([]int{1}, 2)
	if ret != 2 {
		t.Errorf("Expect 0 but received %d", ret)
	}

	ret = CoinChange([]int{1, 2, 5}, 11)
	if ret != 3 {
		t.Errorf("Expect 0 but received %d", ret)
	}

	ret = CoinChange([]int{2}, 3)
	if ret != -1 {
		t.Errorf("Expect 0 but received %d", ret)
	}

}
