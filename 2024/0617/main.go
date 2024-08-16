package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hopewell-energybox_480"

	arr := strings.Split(str, "_")

	fmt.Println(arr)

	//str = "one,two,three,four,five"
	separator := "_"

	// 找到最后一个分隔符的位置
	index := strings.LastIndex(str, separator)
	if index < 0 {
		fmt.Println("Separator not found")
		return
	}

	// 使用 Substring 函数从最后一个分隔符处拆分字符串
	firstPart := str[:index]
	secondPart := str[index+1:]

	fmt.Printf("First part: %s\n", firstPart)
	fmt.Printf("Second part: %s\n", secondPart)

	// 如果需要更一般的拆分，可以继续使用 strings.Split
	//parts := strings.Split(secondPart, separator)
	//fmt.Println("Splitted parts:", parts)
}
