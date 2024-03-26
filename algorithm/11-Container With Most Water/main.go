package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-01-06 16:46
* Email: 364021318@qq.com
*
 */

func maxArea(height []int) int {
	left, right := 0, len(height) - 1
	res := 0
	for left < right {
		// [left, right] 之间的矩形面积
		curArea := func() int {
			if height[left] < height[right] {
				return height[left] * (right - left)
			}
			return height[right] * (right - left)
		}()
		res = func() int {
			if curArea > res {
				return curArea
			}
			return res
		}()
		// 双指针技巧， 移动较低的一边
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	res := maxArea(height)

	fmt.Println(res)
}