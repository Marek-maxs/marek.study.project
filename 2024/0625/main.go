package main

import (
	"fmt"
	"time"
)

func main() {
	//va := 3.98
	//result := ((va - 12420*(va/(va+1))) - (va - 86.4*(va/(va+1))) - (va - 432*(va/(va+1))) - (va - 432*(va/(va+1))) - (va - 138.24*(va/(va+1))) - (va - 345.6*(va/(va+1))) - (va - 345.6*(va/(va+1))) - (va - 345.6*(va/(va+1))) - (va - 172.8*(va/(va+1))) - (va - 544.32*(va/(va+1))) - (va - 54.432*(va/(va+1))) - (va - 496.8*(va/(va+1))) - (va - 496.8*(va/(va+1)))) / 60
	//fmt.Println((va - (345.6 * va / (va + 1))) / 60)
	//fmt.Println(result)
	//fmt.Println(((2.08 - (54.432 * 2.08 / (2.08 + 1)) + 2.08 - (496.8 * 2.08 / (2.08 + 1))) / 60))
	//// 初始化随机数种子
	//rand.Seed(time.Now().UnixNano())
	//
	//// 设定开始值、结束值和总分钟数
	//startValue := 0.0
	//totalValue := 0.0088
	//totalMinutes := 1440
	//
	//// 初始化每分钟的值数组
	//values := make([]float64, totalMinutes)
	//remainingValue := totalValue
	//
	//// 初始化第一个值
	//values[0] = startValue
	//remainingValue -= startValue
	//
	//// 遍历剩余的分钟
	//for i := 1; i < totalMinutes; i++ {
	//	// 计算一个随机的增量，这里使用剩余值的某个比例加上一个随机数
	//	increment := remainingValue / float64(totalMinutes-i)       // 平均分配剩余值
	//	randomIncrement := increment + rand.Float64()*increment*0.1 // 加上一个随机扰动
	//
	//	// 如果加上这个增量会超过剩余值，则使用剩余值作为增量
	//	if remainingValue < randomIncrement {
	//		randomIncrement = remainingValue
	//	}
	//
	//	// 更新当前分钟的值和剩余值
	//	values[i] = values[i-1] + randomIncrement
	//	remainingValue -= randomIncrement
	//
	//	// 为了防止负值，确保剩余值不小于0
	//	if remainingValue < 0 {
	//		remainingValue = 0
	//	}
	//}
	//
	//// 如果最后一分钟的值仍然小于结束值，我们可以将其设置为结束值（这会稍微破坏不规则性）
	//if values[totalMinutes-1] < totalValue {
	//	values[totalMinutes-1] = totalValue
	//}
	//
	//// 遍历并打印每分钟的值
	//for i, value := range values {
	//	fmt.Printf("Minute %d: Value = %.8f\n", i, value)
	//}

	tt := time.Now().Format("2006-01")
	fmt.Println(tt)
}
