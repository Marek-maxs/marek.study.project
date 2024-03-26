package main

import "fmt"

var resume chan int
var yield chan int
var closeFalg chan bool
func integers() chan int {
	yield = make(chan int)
	count := 1

	go func() {
		defer close(yield)
		for {
			select {
			case <-closeFalg:
				fmt.Println("close chan")
				return
			default:
				yield <- count // 这个通道用for循环向通道写值
				// 如果通道里的数据没有被读取, 则不会写下一个数据
				count++
			}

		}
	}()

	return yield
}

func generateInteger() int {
	return <-resume
}

func GetResume() {
	closeFalg = make(chan bool, 1)
	resume = integers() // 将生成整数的通道赋值给resume
	// 调用generateInterger; 就是读取通道
	fmt.Println(generateInteger())
	closeFalg <- true
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}