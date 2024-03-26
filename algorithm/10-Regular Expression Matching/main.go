package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-01-06 15:48
* Email: 364021318@qq.com
*
 */

func isMatch(s string, p string) bool {
	// 备记录
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(p))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 指针 i, j 从索引 0 开始移动
	return dp(s, 0, p, 0, memo)
}

/* 计算p[j..] 是否匹配 s[i..] */
func dp(s string, i int, p string, j int, memo [][]int) bool  {
	m, n := len(s), len(p)
	// base case
	if j == n {
		return i == m
	}
	if i == m {
		if (n - j)%2 == 1 {
			return false
		}
		for ; j + 1 < n; j += 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}

	// 查备忘录，防止重复计算
	if memo[i][j] != -1 {
		return memo[i][j] == 1
	}

	var res bool

	if s[i] == p[j] || p[j] == '.' {
		if j < n-1 && p[j+1] == '*' {
			res = dp(s, i, p, j+2, memo) || dp(s, i+1, p, j, memo)
		} else {
			res = dp(s, i+1, p, j+1, memo)
		}
	} else {
		if j < n-1 && p[j+1] == '*' {
			res = dp(s, i, p, j+2, memo)
		} else {
			res = false
		}
	}
	// 将当前结果记入备忘录
	memo[i][j] = 0
	if res {
		memo[i][j] = 1
	}

	return res
}

func main() {
	res := isMatch("aa", "a")
	fmt.Println(res)
}