package main

// walk path
var ret [][]int
var stack map[int]bool

func travel(nums []int, s map[int]bool) {
	if len(nums) == len(s) {
		ts := make([]int, 0)
		for k, _ := range s {
			ts = append(ts, k)
		}
		ret = append(ret, ts)
		return
	}

	for _, v := range nums {
		if _, ok := stack[v]; ok {
			continue
		}
		stack[v] = true
		travel(nums, stack)
		// back one step
		delete(stack, v)
	}

}

func PermuteTravel(nums []int) [][]int {
	ret = make([][]int, 0)
	stack = make(map[int]bool)

	travel(nums, stack)
	return ret

}

//////////////////////////recursion/////////////////////////////////////////////
var ret [][]int

func PermuteRecursion(nums []int) [][]int {
	if len(nums) == 1 {
		ret := make([][]int, 0)
		ret = append(ret, nums)
		return ret
	}

	ret := make([][]int, 0)
	for k, v := range nums {
		a := []int{}
		a = append(a, nums[:k]...)
		a = append(a, nums[k+1:]...)
		child := Permute(a)
		for _, cv := range child {
			ctmp := []int{v}
			ctmp = append(ctmp, cv...)
			ret = append(ret, ctmp)
		}
	}
	return ret
}
