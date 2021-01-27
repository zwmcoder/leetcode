package main

func MaxArea(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	area := 0
	h := 0
	tmpArea := 0
	for i := length - 1; i >= 0; i-- {
		for j := 0; j <= i ; j++ {
			if height[i] >= height[j] {
				h = height[j]
			} else {
				h = height[i]
			}
			tmpArea = h * (i - j)
			if tmpArea > area {
				area = tmpArea
			}
		}
	}
	return area
}
