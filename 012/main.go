package main

func IntToRoman(num int) string {
	T := []string{"", "M", "MM", "MMM"}
	H := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	TE := []string{"", "X", "XX", "XXX", "XL","L", "LX", "LXX", "LXXX", "XC"}
	U := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return T[num/1000] + H[num%1000/100] + TE[num%100/10] + U[num%10]
}
