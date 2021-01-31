package main

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			n := nums[i] + nums[j] + nums[k]
			if n == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				l := j
				for l < k && nums[l] == nums[j] {
					l++
				}
				j = l
			} else if n > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
