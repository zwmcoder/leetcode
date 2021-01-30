package main

func findCommonPrefix(str1, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}

	var k int
	length := len(str1)
	for k = 0; k < length; k++ {
		if str1[k] == str2[k] {
			continue
		} else {
			break
		}
	}


	ret := str1[:k]
	return ret
}

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	
	common := strs[0]
	for i:=1; i < length; i++ {
		common = findCommonPrefix(common, strs[i])
	}
	return common
}
