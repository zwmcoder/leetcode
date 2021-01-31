package main

import (
//	"fmt"
//	"sort"
)

func getCharacterFromNumber(n string) []string {
	switch n {
	case "2":
		return []string{"a", "b", "c"}
	case "3":
		return []string{"d", "e", "f"}
	case "4":
		return []string{"g", "h", "i"}
	case "5":
		return []string{"j", "k", "l"}
	case "6":
		return []string{"m", "n", "o"}
	case "7":
		return []string{"p", "q", "r", "s"}
	case "8":
		return []string{"t", "u", "v"}
	case "9":
		return []string{"w", "x", "y", "z"}
	}

	return []string{}
}

// This is my first meathod , recursion
func LetterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	if len(digits) == 1 {
		return getCharacterFromNumber(digits)
	}

	length := len(digits)
	sub := LetterCombinations(digits[:length-1])
	str := getCharacterFromNumber(string(digits[length-1]))

	ret := make([]string, 0)
	for _, v := range sub {
		for _, v1 := range str {
			ret = append(ret, v+v1)
		}
	}
	return ret
}

// This is iteration method
func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	dict := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// The result will append all the time, and is pointer, will change
	result = []string{""}
	for i := range digits {
		st := dict[digits[i]] // get mapped strings
		var next []string
		for j := range st {
			c := st[j]
			for _, r := range result {
				next = append(next, r+string(c))
			}
		}

		result = next
	}
	return result
}
