package main

var ret [][]int{{}}
var stack map[int]bool


func travel(nums []int, s map[int]bool) {
	if len(nums) == len(s) {
		ts := make([]int, 0)
		for k, _ := range s {
			ts = append(ts, k)
		}
		ret = append(ret, ts)
	}
	
	for k, v := range nums {
		if _, ok := stack[v]; ok {
			continue
		}
		stack[v] = true
		travel(nums, stack)
		delete(stack, v)
	}

}

func Permute(nums []int) [][]int {
	ret = make([][]int, 0)
	for k, v := range nums {
		if _, ok := stack[v]; ok {
			continue
		}
		stack[v] = true


	}

    
}