package leetcode

import (
	"fmt"
)

func maxArea(height []int) int {

	begin := 0
	end := len(height)-1
	maxAreas := 0
	min := 0
	if len(height) < 2 {
		return 0
	}
	for ;begin < end; {
		min = minInt(height[begin], height[end])
		if (min * (end - begin)) > maxAreas {
			maxAreas = min * (end - begin)
		}
		if height[begin] > height[end] {
			end = end - 1
		} else {
			begin = begin + 1
		}
	}
	return maxAreas
}

func minInt(x, y int) (ret int){

	if x >= y {
		ret = y
	}else{
		ret = x
	}
	return
}

func main() {
	s :=[] int {1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(s))
}