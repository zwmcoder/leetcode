package main

// Use dp
func GenerateParenthesis(n int) []string {
	state := make(map[int][]string)

	state[0] = []string{""}
	state[1] = []string{"()"}

}
