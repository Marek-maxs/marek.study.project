package main

import "fmt"

/**
*
* Author: Marek
* Date: 2023-09-28 23:16
* Email: 364021318@qq.com
*
 */

func lengthOflengestSubstring(s string) int {
	window := make(map[byte]int)

	left, right := 0, 0
	res := 0 // 记录结果
	for right < len(s) {
		c := s[right]
		right++
		// 判断左侧窗口是否收缩
		for window[c] > 1 {
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			window[d]--
		}
		// 在这里更新答案
		res = max(res, right - left)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	str := "pwwkew"
	res := lengthOflengestSubstring(str)
	fmt.Println(res)
}