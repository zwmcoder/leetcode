package main

func IsMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	state := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		d := make([]bool, pLen+1)
		state[i] = d
	}

	state[0][0] = true

	// There have case s is empty "" and p match it, that's the only reason we need
	// initialize state[0][] specially
	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			// For example, s = "",  p = "xxa*", j=4, p[j-1] == '*'
			// The only way is * means zero match, that means "xx" match the s
			// that is state[0][j] = state[0][j-2]
			state[0][j] = state[0][j-2]
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				state[i][j] = state[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					// case 1: This is the case * means zero,  s = "aabc" p = "aabck*"
					state[i][j] = state[i][j-2]
				} else {
					// For case 2 and case 3 s[i-1]==p[j-2]=='c'
					// case 2: This is the case * means 1 repeat,  s = "aabc" p = "aabc*"  : state[i][j-1]
					// case 3: This is the case * means 2 repeat,  s = "aabcc" p = "aabc*" : state[i-1][j]
					// case 4: This is the case * means 0 repeat,  s = "aabcc" p = "aabcc.*" : state[i][j-2]
					state[i][j] = state[i][j-1] || state[i-1][j] || state[i][j-2]
				}
			}
		}
	}
	return state[sLen][pLen]
}
