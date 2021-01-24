package main

func IsPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	ret := 0
	for x > ret {
		ret = ret*10 + x%10
		x = x / 10
	}

	if ret == x || x == ret/10 {
		return true
	}
	return false
}
