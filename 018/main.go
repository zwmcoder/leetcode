package main

//	"fmt"
import "sort"

func FourSum(nums []int, target int) [][]int {
	var result [][]int

	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length-3; i++ {
		// skip duplicate
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			k := j + 1
			q := length - 1

			for k < q {
				n := nums[i] + nums[j] + nums[k] + nums[q]
				if n == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[q]})
					// skip duplicate k
					l := k
					for l < q && nums[l] == nums[k] {
						l++
					}
					k = l
				} else if n > target {
					q--
				} else {
					k++
				}
			}
		}
	}
	return result
}
