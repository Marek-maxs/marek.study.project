package main

import (
	"sync"
	"time"
)

// 测试执行一个很大 groutines 数量的程序，看看效果是不是比预期的要优秀
func main() {
	var wg sync.WaitGroup
	wg.Add(100000)

	for i := 0; i < 100000; i++ {
		go func() {
			time.Sleep(1 * time.Minute)
		}()
	}

	wg.Wait()
}
