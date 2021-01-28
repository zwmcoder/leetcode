package main

func getInt(v string) int {
	switch v {
		case "I":
			return 1
		case "V":
			return 5
		case "X":
			return 10
		case "L":
			return 50
		case "C":
			return 100
		case "D":
			return 500
		case "M":
			return 1000
	}
	return 0
}

func RomanToInt(s string) int {
	length := len(s)

	ret := 0
	skip := false
	for k := range s {
		if skip {
			// The previous LOOP already count this character like IV for V.
			skip = false 
			continue
		}
		
		if k == length - 1 {
			ret = ret + getInt(string(s[k]))
			break
		}
		if getInt(string(s[k])) < getInt(string(s[k+1])) {
				ret = ret + getInt(string(s[k+1])) - getInt(string(s[k]))
				skip = true
		} else {
				ret = ret + getInt(string(s[k]))
		}
	}
	return ret
}
