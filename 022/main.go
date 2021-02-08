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
