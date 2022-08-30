package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	var wg sync.WaitGroup
	//give a fixed value
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// 可以保证即使在多核系统下 count++ 也是一个原子操作
			// 对 addr 指向的值加上 delta. 如果 将 delta 设置成负值，加法就变成了减法
			atomic.AddInt64(&count, 1)
			// continuous reduction after implementation
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&count))

	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&count))

	studyAtomic()
}

type T struct {
	x int16
	y int64
}

// 在现在的操作系统中，写的地址基本是对齐的。 32位系统中，变量的起始地址都是4倍速，64位系统中，变量的起始
// 地址都是8的倍数。如果在32位的系统上进行64位的写操作，系统可能需要两个指令才能完成（你还记得GO并发任务编排利器之WaitGroup
// 中 state 字段会对不同系统的不同处理吗？）。 对齐地址的读写，不会导致其它协程只看到写了一半的数据。
// atomic 包提供的方法会提供内存屏障的功能，所以，atomic 不仅仅可以保证赋值的数据完整性，还能保证数据的可见性，一旦一个核
// 更新了该地址的值，其它处理器总是能读取到它的最新值。
func studyAtomic() {
	var v atomic.Value
	var t = T{}
	v.Store(t)
	a := T{x: 1, y: 2}
	b := T{x: 0, y: 0}
	go func() {
		for {
			go func() {
				t = a
				v.Store(t)
			}()
			go func() {
				t = b
				v.Store(t)
			}()
		}
	}()
	for {
		fmt.Println(v.Load().(T))
	}
}
