package main

import "fmt"

func isPalindrome(x int) bool {

	if x == 0 {
		return true
	}else if x < 0 {
		return  false
	}else {
		var ret int = 0
		c_x := x
		for ;c_x != 0; {
			ret = ret + c_x % 10
			c_x = c_x / 10
			if c_x != 0 {
				ret = ret * 10
			}
		}
		if ret == x {
			return true
		} else {
			return false
		}
	}
}

func main() {

	fmt.Println(isPalindrome(323))
}