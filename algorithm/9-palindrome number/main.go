package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-01-06 14:56
* Email: 364021318@qq.com
*
 */

// IsPalindrome 是判断一个整数是否为回文数的函数
func IsPalindrome(x int) bool {
	// 如果 x 是负责，那么它不可能为回文数， 直接返回 false
	if x < 0 {
		return false
	}

	// temp 是 x 的副本
	temp := x
	// y 是 x 翻转后的数字
	y := 0

	// 将temp逆序生成 y
	for temp > 0 {
		lastNum := temp % 10
		temp = temp / 10
		y = y * 10 + lastNum
	}
	// 如果 x 和 y 相等， 那么 x 就是回文数
	return y == x
}

func main() {
	res := IsPalindrome(10)
	fmt.Println(res)
}