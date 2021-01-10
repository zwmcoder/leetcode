package main

func Convert(s string, numRows int) string {
	if numRows == 1 || len(s) == 1 {
		return s
	}
	length := len(s)
	cList := make(map[int][]byte, numRows)

	for i := 0; i < length; i++ {
		isEven := (i / (numRows - 1)) & 0x1
		index := 0
		if isEven == 0 {
			// For the example 2, it refer to the column P, A, Y  and I, S, H(even number column)
			index = i % (numRows - 1)
		} else {
			// For the example 2, it refer to the slash P, A, L and I, R, I(odd numver slash)
			index = numRows - 1 - i%(numRows-1)
		}
		cList[index] = append(cList[index], s[i])
	}

	ret := make([]byte, 0)
	for k := 0; k < numRows; k++ {
		ret = append(ret, cList[k]...)
	}

	return string(ret)
}
