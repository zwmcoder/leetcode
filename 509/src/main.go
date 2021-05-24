package main

func Fib(n int) int {
	var tmp1, tmp2, ret int
	tmp2 = 1

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for idx := 0; idx < n-1; idx++ {
		ret = tmp1 + tmp2
		tmp1 = tmp2
		tmp2 = ret
	}
	return ret
}
