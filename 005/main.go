package main

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return string(s[0])
	}
	retstring := ""
	maxlen := 0
	for k := 1; k < len(s); k++ {
		var lengthodd, lengtheven, length int
		var submaxodd, submaxeven, submax string
		// For the odd string
		submaxodd = string(s[k])
		submaxeven = string(s[k])
		lengthodd = 1
		lengtheven = 1
		for i := 1; ; i++ {
			if k-i < 0 || k+i >= len(s) {
				break
			}
			if s[k-i] == s[k+i] {
				submaxodd = s[k-i:k+i] + string(s[k+i])
				lengthodd = 2*i + 1
			} else {
				break
			}
		}

		// For the even
		for i := 1; ; i++ {
			if k-i < 0 || k+i-1 >= len(s) {
				break
			}
			if s[k-i] == s[k+i-1] {

				submaxeven = s[k-i:k+i-1] + string(s[k+i-1])
				lengtheven = 2 * i
			} else {
				break
			}
		}

		if lengthodd >= lengtheven {
			submax = submaxodd
			length = lengthodd
		} else {
			submax = submaxeven
			length = lengtheven
		}

		if length > maxlen {
			retstring = submax
			maxlen = length
		}
	}
	return retstring
}
