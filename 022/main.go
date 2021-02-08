package main

// Golang recursion
func GenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	} else if n == 1 {
		return []string{"()"}
	}

	ret := make([]string, 0)
	for i := 0; i < n; i++ {
		left := GenerateParenthesis(i)
		right := GenerateParenthesis(n - 1 - i)
		for _, lv := range left {
			for _, rv := range right {
				ret = append(ret, "("+lv+")"+rv)
			}
		}
	}
	return ret
}

// DP solution
func GenerateParenthesisDP(n int) []string {
	ret := make(map[int][]string)

	ret[0] = []string{""}
	ret[1] = []string{"()"}

	for idx := 2; idx <= n; idx++ {
		for i := 0; i < idx; i++ {
			for _, lv := range ret[i] {
				for _, rv := range ret[idx-1-i] {
					ret[idx] = append(ret[idx], "("+lv+")"+rv)
				}
			}
		}
	}
	return ret[n]
}
