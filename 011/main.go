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


// use two pointer, one at the start and one at the end. Each time move the pointer
// which is smaller.
// If move the pointer which point the bigger one, of cause the area will become smaller, no matter the
// next value is bigger or smaller
func MaxAreaOn(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	area := 0
	tmpArea := 0 
	i := 0
	j := length - 1 

  for ; i < j ; {
		if height[i] <= height[j] {
			tmpArea = height[i] * (j-i)
			i++
		} else {
			tmpArea = height[j] * (j-i)
			j--
		}
		if tmpArea > area {
			area = tmpArea
		}
	} 

	return area
}
