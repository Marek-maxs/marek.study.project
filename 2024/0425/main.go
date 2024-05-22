package main

import (
	"encoding/binary"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// 假设有一个int类型的变量
	var num int = 2450001

	// 使用strconv.FormatInt将int转换为16进制字符串，并指定基数为16
	hexStr := strconv.FormatInt(int64(num), 16)

	// 输出结果，注意默认是小写
	fmt.Println("16进制表示:", hexStr)

	// 十六进制字符串
	hexStr = "0004b4"

	// 使用strconv.ParseInt将十六进制字符串转换为十进制整数
	// 第二个参数指定了进制为16（十六进制）
	// 第三个参数指定了结果的大小（在这个例子中为int64）
	decimal, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		// 处理错误
		fmt.Println("转换错误:", err)
		return
	}

	// 输出十进制结果
	fmt.Println("十进制表示:", decimal)

	datas := []byte{0x00, 0xc7, 0x63}

	fmt.Println(binary.BigEndian.Uint16(datas))

	fmt.Println(GetExpressionDatapointID("abs({id_3245}+{id_3247})"))
}


func GetExpressionDatapointID(expression string) [][]string {
	rex := regexp.MustCompile(`\{\w+_(\d+)\}`)
	expressionArr := rex.FindAllStringSubmatch(expression, -1)

	return expressionArr
}