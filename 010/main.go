package main

func IsMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	dp := make([][]bool, sLen + 1)
	for i:= 0; i <= sLen; i++ {
		d := make([]bool, pLen + 1)
		dp[i] = d
	}

	dp[0][0] = true

	
	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			// For example, s = "",  p = "xxa*", j=4, p[j-1] == '*'
			// The only way is * means zero match, that means "xx" match the s
			// that is dp[0][j] = dp[0][j-2]
			dp[0][j] = dp[0][j-2]
		}

	}



}
