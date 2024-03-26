package main

import "strings"

/**
*
* Author: Marek
* Date: 2024-01-06 19:07
* Email: 364021318@qq.com
*
 */


func letterCombinations(digits string) []string {
	mapping := []string {
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	var res []string
	if len(digits) == 0 {
		return res
	}
	// 从digits[0]开始进行回溯
	backtrack(digits, 0, &strings.Builder{}, mapping, &res)
	return res
}

// 回溯算法主函数
func backtrack(digits string, start int, sb *strings.Builder, mapping []string, res *[]string)  {
	if sb.Len() == len(digits) {
		// 到达回溯树底部
		*res = append(*res, sb.String())
		return
	}
	// 回溯算法框架
	for i := start; i < len(digits); i++ {
		digit := digits[i] - '0'
		for _, c := range mapping[digit] {
			// 做选择
			sb.WriteRune(c)
			// 递归下一层回溯
			backtrack(digits, i+1, sb, mapping, res)
			// 撤销选择
			sb.Truncate(sb.Len() -1)
		}
	}
}