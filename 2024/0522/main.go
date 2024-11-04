package main

import (
	"fmt"
	"time"
)

/*
在这个例子中，select会阻塞，直到interrupt changel 有数据可读。
一旦接收到数据，select 就会结束，程序执行后续的关闭操作。

然而，在大多数并发场景中，select 与 for 循环结合使用，以便在多个channel 之间持续轮询，直到满足某种退出条件。
在两个或更多goroutine 之间使用select时，外层的for循环通常是用来处理以下情况：
1 持久监听： select 可能会持续等待来自不同goroutine 的消息，这意味着我们需要保持select 语句的活性，直到遇到某个特定的退出条件。
for 循环可以保证这一点，直到出现特定的退出条件（例如，所有的channel 都被关闭，或者接收到特定的信号）。
2 非阻塞性检查： 即使没有㶢可读或可写，for循环也可以配合default 子句，用于周期性地检查某些条件，或者执行其他的非阻塞操作。
3 控制并发行为： 通过for 循环， 我们可以控制并发行为，例如限制并发的数量，或者在处理完一批任务后才启动。
4 处理不确定的结束条件： 在并发环境中， 何时结束往往不是预先确定，for 循环允许我们持续监控直到满足条件， 比如所有的工作都被完成。

*/

func main() {
	intChan1 := make(chan int)
	intChan2 := make(chan int)

	// 启动两个goroutines, 分别向两个channel 发送数据
	go func() {
		for i := 1; i <= 5; i++ {
			intChan1 <- i
			time.Sleep(100 * time.Microsecond)
		}
		close(intChan1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			intChan2 <- i
			time.Sleep(150 * time.Microsecond)
		}
		close(intChan2)
	}()

	// 使用for循环处理两个channel的数据，直到它们都关闭
	for {
		select {
		case value := <-intChan1:
			fmt.Printf("Received from channel 1: %d\n", value)
		case value := <-intChan2:
			fmt.Printf("Receive from channel 2: %d\n", value)
		case <-time.After(1 * time.Second):
			fmt.Println("Both channels closed, exiting.")
			return
		}
	}
}
