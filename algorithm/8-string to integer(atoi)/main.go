package main

import (
	"fmt"
	"math"
)

func main() {
	str := "string"
	res := MyAtoi(str)

	fmt.Println(res)
}

func MyAtoi(str string) int {
	n := len(str)
	i := 0
	// 记录正负号
	sign := 1
	// 用 long 避免 int 溢出
	var res int64 = 0
	// 跳过前导空格
	for i < n && str[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}
	// 记录符号位
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}
	if i == n {
		return 0
	}
	// 统计数字位
	for i < n && '0' <= str[i] && str[i] <= '9' {
		res = res * 10 + int64(str[i]-'0')
		if res > math.MaxInt32 {
			break
		}
		i++
	}
	// 如果溢出，强转成 int 就会和真实值不同
	if res > math.MaxInt32 {
		if sign == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	return int(res) * sign
}