package main

func Reverse(x int) int {
	negative := false 
	if x < 0 {
		negative = true
		x = -x
	}
	ret := 0

	if x > 0x7FFFFFFF {
		return 0
	}

	for;x > 0; {
		tmp := x%10
		ret = ret * 10 + tmp
		x = x / 10
	}
	   
	if ret > 0x7FFFFFFF{
		return 0
	}
	
	if negative {
		ret = -ret
	}

	return ret
}
