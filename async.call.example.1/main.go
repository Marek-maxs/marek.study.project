package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func AsyncCallOne() {
	// init context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond *100))
	// last close ctx
	defer cancel()
	// create async goroution
	go func(ctx context.Context) {
		// program
	}(ctx)

	// listen channel and ctx
	select {
	case <-ctx.Done():
		fmt.Println("call successfully!!")
		return
	case <-time.After(time.Duration(time.Millisecond * 10)):
		fmt.Println("timeout!!")
		return
	}
}

func AsyncCallTow() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	timer := time.NewTimer(time.Duration(time.Millisecond * 800))

	go func(ctx context.Context) {
		// program
	}(ctx)

	select {
	case <- ctx.Done():
		timer.Stop()
		timer.Reset(time.Second)
		fmt.Println("call successfully!!")
		return
	case <-timer.C:
		fmt.Println("timeout!!")
		return
	}
}

func AsynCallThree() {
	ctx := context.Background()
	done := make(chan struct{}, 1)

	go func(ctx context.Context) {
		done <- struct{}{}
	}(ctx)
	select {
	case <-done:
		fmt.Println("call successfull!!!")
		return
	case <-time.After(time.Duration(800 * time.Millisecond)):
		fmt.Println("timeout!!!")
		return
	}
}


// 课后练习 1.2
// 基于 Channel 编写一个简单的单协程生产消费者模型
// 要求如下你：
// 1. 队列： 队列长度10 ， 队列 元素类型为 int
// 2. 生产者： 每1 秒往队列中放入一个类型为 int 的元素， 队列满时生产者可以阻塞
// 3. 消费者： 每 2 秒从队列中获取一个元素并打印，队列为空时消费者阻塞
// 4. 主协程 30 秒后要求所有子协程退出。
// 5. 要求优雅退出， 即协程退出前，要先消费完所有 int.

// 知识 点
// 1. 切片 的零值 也是可用的。
// 2. context.WithTimeout

var (
	wg sync.WaitGroup
	p Producer
	c Consumer
)

type Producer struct {
	Time int
	Interval int
}

type Consumer struct {
	Producer
}

func (p Producer) produce(queue chan<- int, ctx context.Context) {
	go func() {
		LOOP:
			for {
				p.Time = p.Time + 1
				queue <- p.Time
				fmt.Printf("生产者进行第%d次生产， 值： %d\n", p.Time, p.Time)
				time.Sleep(time.Duration(p.Interval) * time.Second)

				select {
				case <-ctx.Done():
					close(queue)
					break LOOP
				}
			}
			wg.Done()
	}()
}

func (c Consumer) consume(queue <-chan int, ctx context.Context) {
	go func() {
		LOOP:
			for {
				c.Time++
				val := <-queue
				fmt.Printf("-->消费者进行第%d次消费， 值: %d\n", c.Time, val)
				time.Sleep(time.Duration(c.Interval) * time.Second)

				select {
				case <-ctx.Done():
					var remains []int // 知识点： 切片的零值也是可用的
					for val = range queue {
						remains = append(remains, val)
						fmt.Printf("-->消费者: 最后一次消费， 值为: %v \n", remains)
						break LOOP
					}
				}
			}
			wg.Done()
	}()
}

func main() {
	wg.Add(2)

	// 知识点： context.Timeout
	timeout := 30
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	queue := make(chan int, 10)

	p.produce(queue, ctx)
	c.consume(queue, ctx)
	fmt.Println("main waiting...")
	wg.Wait()
	fmt.Println("done")
}

// 启动命令
// go run main/main.go -m wb
// go run main/main.go -m je

func init() {
	// 解析程序入参， 运行模式
	mode := flag.String("m", "wb", "请输入运行模式: \nwb (温饱模式) 生产速度快过消费速度、\nje (饥饿模式) 生产速度慢于消费速度")
	flag.Parse()

	p = Producer{}
	c = Consumer{}

	if *mode == "wb" {
		fmt.Println("运行模式： wb (温饱模式)生产速度快过消费速度")
		p.Interval = 1 // 每隔 1 秒生产一次
		c.Interval = 5 // 每隔5秒 消费一次

	} else {
		fmt.Println("运行模式： je (饥饿模式)生产速度慢于消费速度")
		p.Interval = 5 // 每隔 5 秒 生产一次
		c.Interval = 1 // 每隔 1 秒 消费一次
	}
}
