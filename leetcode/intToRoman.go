package leetcode

func intToRoman(num int) string {

	var ret []byte
	num_c := num
	num_p := 0
	for ;num_c > 0; {
		if num_c >= 1000 {
			//
			num_p = num_c / 1000
			for i:=0; i<num_p; i++ {
				ret = append(ret, 'M')
			}
			num_c = num_c % 1000
		}else if (num_c >= 100) && (num_c < 1000) {
			//
			num_p = num_c / 100
			switch num_p{
			case 1, 2, 3:
				for i:=0; i<num_p; i++ {
					ret = append(ret, 'C')
				}
			case 4:
				ret = append(ret, 'C', 'D')
			case 5:
				ret = append(ret, 'D')
			case 6, 7, 8:
				ret = append(ret, 'D')
				for i:=0; i<num_p-5; i++ {
					ret = append(ret, 'C')
				}
			case 9:
				ret = append(ret, 'C', 'M')
			}
			num_c = num_c % 100
		}else if (num_c >= 10) && (num_c < 100) {
			//
			num_p = num_c / 10
			switch num_p{
			case 1, 2, 3:
				for i:=0; i<num_p; i++ {
					ret = append(ret, 'X')
				}
			case 4:
				ret = append(ret, 'X', 'L')
			case 5:
				ret = append(ret, 'L')
			case 6, 7, 8:
				ret = append(ret, 'L')
				for i:=0; i<num_p-5; i++ {
					ret = append(ret, 'X')
				}
			case 9:
				ret = append(ret, 'X', 'C')
			}
			num_c = num_c % 10
		}else if (num_c >= 1) && (num_c < 10) {
			//
			num_p = num_c
			switch num_p{
			case 1, 2, 3:
				for i:=0; i<num_p; i++ {
					ret = append(ret, 'I')
				}
			case 4:
				ret = append(ret, 'I', 'V')
			case 5:
				ret = append(ret, 'V')
			case 6, 7, 8:
				ret = append(ret, 'V')
				for i:=0; i<num_p-5; i++ {
					ret = append(ret, 'I')
				}
			case 9:
				ret = append(ret, 'I', 'X')
			}
			num_c = num_c / 10
		}
	}
	return string(ret)
}
