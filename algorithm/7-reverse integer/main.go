package main

import "math"

/**
*
* Author: Marek
* Date: 2023-10-05 22:14
* Email: 364021318@qq.com
*
 */

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}