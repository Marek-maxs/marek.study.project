package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isValidFloat(s string) (float64, bool) {
	// 去除字符串两端的空白字符
	s = strings.TrimSpace(s)

	// 尝试将字符串解析为浮点数
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		// 解析错误，不是有效的浮点数
		return 0, false
	}

	// 检查是否为无限值
	if math.IsInf(f, 0) {
		fmt.Println(f)
		// 是无限值，不是有限的数值
		return 0, false
	}

	// 是有效的有限数值
	return f, true
}

func main() {
	// 示例输入值
	values := []string{"123.45", "NaN", "Infinity", "abc", ""}

	for _, value := range values {
		if f, ok := isValidFloat(value); ok {
			fmt.Printf("Valid float: %f\n", f)
		} else {
			fmt.Printf("Invalid float: %s\n", value)
		}
	}
}
