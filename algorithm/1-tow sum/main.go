package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 维护 val -> index 的映射
	valToIndex := make(map[int]int)

	for i, num := range nums{
		// 查表，看看是否有能和 nums[i] 凑出 target 的元素
		need := target - num
		if valToIndex[need] != 0 {
			return []int{valToIndex[need] - 1, i}
		}
		// 存入 val -> index 的映射
		valToIndex[num] = i + 1
	}

	return nil
}

func main()  {
	nums := []int{-1,-2,-3,-4,-5}
	target :=-8
	// 从数组中找到两个元素和和 满足 目标值 的数据，返回 数组对应的键名
	res  := twoSum(nums, target)
	// result: [2 4]
	fmt.Println("result:", res)
}