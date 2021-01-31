package main

import (
	"fmt"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	fmt.Println(nums)
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
		fmt.Printf("111:%d %d %d\n", i, j, k)
		for j < k {
			n = nums[i] + nums[j] + nums[k]
			fmt.Printf("%d %d %d\n", i, j, k)
			fmt.Printf("n %d\n", n)
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
	fmt.Println("The result is ", n)
	return retn[ret]
}
