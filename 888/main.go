package main

//	"fmt"
func FairCandySwap(A []int, B []int) []int {
	var sumA, sumB int
	for _, v := range A {
		sumA += v
	}
	for _, v := range B {
		sumB += v
	}

	sub := sumA - sumB

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if 2*(A[i]-B[j]) == sub {
				return []int{A[i], B[j]}
			}
		}
	}
	return []int{}
}
