package leetcode

func reverse(x int32) int32 {

	index := 0
	if x >= 0 {
		index = 0
	}else{
		index = 1
		x = -x
	}
	var res int32
	var mid int32
	for ;x > 0; {
		mid = x % 10
		x = x / 10
		if mid == 0 && res == 0 {
			continue
		}
		res = mid + res
		if x > 0 {
			res = res * 10
		}
	}
	if index == 0{
		return res
	}else{
		return -res
	}

}
