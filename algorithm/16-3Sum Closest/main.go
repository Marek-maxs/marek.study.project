package main

import (
	"fmt"
	"math"
	"sort"
)

/**
*
* Author: Marek
* Date: 2024-01-06 18:32
* Email: 364021318@qq.com
*
 */

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	// 先排序数组
	sort.Ints(nums)
	// 记录三数之和与目标值的偏差
	delta := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		// 固定 nums[i] 为三数之和中的第一个数，
		// 然后对 nums[i+1..] 搜索接近 target - nums[i] 的两数之和
		sum := nums[i] +twoSumClosest(nums, i+1, target-nums[i])
		if int(math.Abs(float64(delta))) > int(math.Abs(float64(target-sum))) {
			delta = target - sum
		}
	}
	return target - delta
}

// 在 nums[start...] 搜索最接近 target 的两数之和
func twoSumClosest(nums []int, start int, target int) int {
	lo, hi := start, len(nums)-1
	// 记录两数之生发目标值的偏差
	delta := math.MaxInt32
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if int(math.Abs(float64(delta))) > int(math.Abs(float64(target-sum))) {
			delta = target - sum
		}
		if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return target - delta
}

func main() {
	nums := []int{-1,2,1,-4}
	target := 1

	res := threeSumClosest(nums, target)
	fmt.Println(res)
}