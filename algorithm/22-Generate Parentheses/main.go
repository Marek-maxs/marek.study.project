package main

/**
*
* Author: Marek
* Date: 2024-01-07 0:51
* Email: 364021318@qq.com
*
 */

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	// 记录所有合法的括号组合
	res := []string{}
	// 回溯过程中的路径
	var track string
	// 可用的左括号和右括号数量初始化为 n
	backtrack(n, n, &track, &res)

	return res
}

// 可用的左括号数量为 left 个， 可用的右括号数量为 right 个
func backtrack(left int, right int, track *string, res *[]string) {
	// 若左路括号剩下的多，说明不合法
	if right < left {
		return
	}
	// 数量小于0 肯定是不合法的
	if left < 0 || right < 0 {
		return
	}
	// 当所有括号都恰好用完时，得到一个合法的括号组合
	if left == 0 && right == 0 {
		*res = append(*res, *track)
		return
	}

	// 尝试放一个左路括号
	*track += "(" // 选择
	backtrack(left-1, right, track, res)
	*track = (*track)[:len(*track)-1]
	// 尝试放一个右括号
	*track += ")"
	backtrack(left, right-1, track, res)
	*track = (*track)[:len(*track)-1] // 撤消选择
}