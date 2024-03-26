package main

import (
	"fmt"
	"sort"
)

/**
*
* Author: Marek
* Date: 2024-01-06 17:33
* Email: 364021318@qq.com
*
 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// n 为 3 ， 从nums[0] 开始计算和为 0 的三元组
	return nSumTarget(nums, 3, 0, 0)
}

// n 填写想求的是几数之和， start 从哪个索引开始计算 （一般填0），target 填想凑出的目标和
func nSumTarget(nums []int, n int, start int, target int) [][]int {
	sz := len(nums)
	var res [][]int
	// 至少是2Sum， 且数组大小应该小于 n
	if n < 2 || sz < n {
		return res
	}
	// 2Sum 是 base case
	if n == 2 {
		// 双指针那一套操作
		lo, hi := start, sz - 1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				for lo < hi && nums[lo] ==  left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		// n > 2 时， 递归计算（n -1） Sum 的结果
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n - 1, i + 1, target - nums[i])
			for _, arr := range sub {
				// (n-1) Sum 加上 nums[i] 就是 nSum
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < sz - 1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	res := threeSum(nums)
	fmt.Println(res)
}