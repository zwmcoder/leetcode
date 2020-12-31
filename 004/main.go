package main

import (
	"fmt"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		return FindMedianSortedArrays(nums2, nums1)
	}

	idxMin := 0
	idxMax := m
	halfOfAll := (m + n + 1) >> 1

	fmt.Printf("%d, %d, %d \n", m, n, halfOfAll)
	var idxM, idxN int
	for idxM <= m { // <== this is the second mistake, at first i use: for idxM < m,  this make Case 2 test fail
		idxM = (idxMin + idxMax) >> 1
		idxN = halfOfAll - idxM
		if idxM > 0 && nums1[idxM-1] > nums2[idxN] {
			fmt.Printf("case1 idxN: %d idxM: %d\n", idxN, idxM)
			// cursor move to left
			idxMax = idxM - 1
		} else if idxM < m && idxN > 0 && nums2[idxN-1] > nums1[idxM] {
			fmt.Printf("case2 idxN: %d idxM: %d\n", idxN, idxM)
			// cursor move to right
			idxMin = idxM + 1
		} else {
			fmt.Printf("case3 idxN: %d idxM: %d\n", idxN, idxM)
			var maxLeft, minRight int
			if idxM == 0 {
				maxLeft = nums2[idxN-1]
			} else if idxN == 0 {
				maxLeft = nums1[idxM-1]
			} else {
				maxLeft = max(nums2[idxN-1], nums1[idxM-1])
			}

			if (m+n)&0x1 == 0x1 { // <== this is the first mistake i miss
				return float64(maxLeft)
			}

			if idxM == m {
				minRight = nums2[idxN]
			} else if idxN == n {
				minRight = nums1[idxM]
			} else {
				minRight = min(nums2[idxN], nums1[idxM])
			}
			fmt.Printf("%d, %d\n", maxLeft, minRight)
			// cursor is the right place
			return (float64(maxLeft) + float64(minRight)) / 2
		}
	}
	fmt.Printf("idxM: %d\n", idxM)
	return 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
