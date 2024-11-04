package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

/**
*
* Author: Marek
* Date: 2024-01-24 15:26
* Email: 364021318@qq.com
*
 */

func main() {
	// 获取当前时间
	now := time.Now()

	// 计算本周第一天的时间点（星期一）
	startOfWeek := now.AddDate(0, 0, int(-int(now.Weekday())+1))

	// 计算本周最后一天的时间点（星期日）
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	// 打印本周的起始时间和结束时间
	fmt.Println("本周的起始时间为：", startOfWeek)
	fmt.Println("本周的结束时间为：", endOfWeek)

	expression := "abs({id_675598}-{id_690487}-({id_675597}*1000)dfwefwe)"
	rex := regexp.MustCompile(`\{\w+_(\d+)\}`)
	expressionArr := rex.FindAllStringSubmatch(expression, -1)
	for _, val := range expressionArr {
		fmt.Println(val[1])
	}
	fmt.Println(expressionArr)

	txt := "abc_123"
	ch, num, _ := strings.Cut(txt, "_")
	fmt.Println(ch, num)
}