package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	ret := Fib(0)
	if ret != 0 {
		t.Errorf("Expect 0 but received %d", ret)
	}

	ret = Fib(1)
	if ret != 1 {
		t.Errorf("Expect 0 but received %d", ret)
	}

	ret = Fib(2)
	if ret != 1 {
		t.Errorf("Expect 0 but received %d", ret)
	}

	ret = Fib(3)
	if ret != 2 {
		t.Errorf("Expect 0 but received %d", ret)
	}

	ret = Fib(4)
	if ret != 3 {
		t.Errorf("Expect 0 but received %d", ret)
	}
}
