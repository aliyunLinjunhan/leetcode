package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(str string) int {

	str = strings.TrimSpace(str)
	var position int
	for index, value := range str {

		if index == 0 {
			if !unicode.IsNumber(value) && value != '+' && value != '-' {
				return 0
			}
		} else {
			if !unicode.IsNumber(value) {
				position = index
				break
			}
		}
	}
	if position != 0 {
		c := []byte(str)
		c = c[0:position]
		fmt.Println(position,c)
		str = string(c)
	}
	ret, _ := strconv.Atoi(str)
	if ret > 2147483647 || ret < -2147483648 {
		if ret > 0 {
			return 2147483647
		} else {
			return -2147483648
		}
	}
	return ret
}

func main() {

	fmt.Println(myAtoi("42"))
}
