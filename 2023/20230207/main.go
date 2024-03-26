package main

import (
	"fmt"
	"time"
)

/**
*
* Author: Marek
* Date: 2023-02-07 9:18
* Email: 364021318@qq.com
*
* channel 通道的练习，有缓存的通道
*
*
 */

func main() {
	//var wg sync.WaitGroup
	ch := make(chan int, 5)
	//wg.Add(11)
	go recv(ch)
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("发送成功", i)
	}
	for {
		if len(ch) == 0 {
			break
		}
	}
	//wg.Wait()
	close(ch)
}

func recv(c chan int) {
	//wg.Done()
	for {
		ret := <-c
		fmt.Println("接收成功", ret)
		time.Sleep(time.Second * 1)
	}
}