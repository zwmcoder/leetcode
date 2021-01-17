package main

func MyAtoi(s string) int {
	ret := 0
	negative := false

	i := 0
	for i = range s {
		if s[i] == ' ' {
				continue
		}
		break
	}
	
	ts := s[i:]
	
	for k := range ts {
		if ts[k] == '-' {
			if k != 0 {
				// If "-" is at the middile
				break
			}
			negative = true
			continue
		} else if ts[k] == '+' {
			if k != 0{
				// if "+" is at the middle
				break
			}
			continue
		}
		if int32(ts[k])-int32('0') < 0  || int32(ts[k]) - int32('9') > 0 {
			break
		}
		ret = ret*10 + int(ts[k] - '0')
		if ret > 0x7FFFFFFF {
			if negative {
				ret = 0x80000000
				break
			} else {
				ret = 0x7FFFFFFF
				break;
			}
		}
	}


	if negative == true {
		ret = -ret
	}

	return ret
}
