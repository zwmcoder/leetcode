package main

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	var i, j, k int
	ret := 0xFFFFFFFF

	retn := make(map[int]int)
	n := 0
	for ; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		j = i + 1
		k = len(nums) - 1
		for j < k {
			n = nums[i] + nums[j] + nums[k]
			if n > target {
				if n-target < ret {
					ret = n - target
					retn[ret] = n
				}
				k--
			} else if n < target {
				if target-n < ret {
					ret = target - n
					retn[ret] = n
				}
				j++
			} else {
				return target
			}

		}
	}
	return retn[ret]
}
