package main

import "fmt"

// GenericTest 定义一个自己的泛型
type GenericTest interface {
	int | int8 | int16 | int64 | float32 | float64 | string
}

func main() {
	fmt.Println(sum(10, 10))
	fmt.Println(sum(10.50, 10.50))
	fmt.Println(sum("hello", ",word"))
}

// 使用泛型
func sum[T GenericTest](a T, b T) T {
	return a + b
}