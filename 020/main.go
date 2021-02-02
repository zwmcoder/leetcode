package main

func isPair(s byte, d byte) bool {
	return (s == '{' && d == '}') ||
		(s == '(' && d == ')') || (s == '[' && d == ']')

}

// The thought is
// Use stack, range string s and push in stack one by one, if match the isPair,  then pop
// out the top
// golang do not have stack, so use array to simulate. However, the bottom of stack may change
// That`s why only when the whole array is empty, then the result is true
func IsValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	stack := make([]byte, len(s))
	stack[0] = s[0]
	i := 0
	for k := 1; k < len(s); k++ {
		if k > 0 && isPair(stack[i], s[k]) {
			stack[i] = 0
			if i > 0 {
				i--
			}
		} else {
			i++
			stack[i] = s[k]
		}
	}

	for _, v := range stack {
		if v != 0 {
			return false
		}
	}
	return true
}
