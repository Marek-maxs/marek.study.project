package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-01-06 17:05
* Email: 364021318@qq.com
*
 */

func intToRoman(num int) string {
	var m = []string{"I", "V", "X", "L", "C", "D", "M", ""}
	var res string
	var n int
	for i := 1; num > 0; i++ {
		n = 2*i - 1
		switch num % 10 {
		case 1:
			res = m[n-1] + res
		case 2:
			res = m[n-1] + m[n-1] + res
		case 3:
			res = m[n-1] + m[n-1] + m[n-1] + res
		case 4:
			res = m[n-1] + m[n] + res
		case 5:
			res = m[n] + res
		case 6:
			res = m[n] + m[n-1] + res
		case 7:
			res = m[n] + m[n-1] + m[n-1] + res
		case 8:
			res = m[n] + m[n-1] + m[n-1] + m[n-1] + res
		case 9:
			res = m[n-1] + m[n+1] + res
		}
		num = num / 10
	}
	return res
}

func main() {
	res := intToRoman(3)

	fmt.Println(res)
}