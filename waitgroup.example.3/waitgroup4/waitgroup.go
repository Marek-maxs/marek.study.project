package waitgroup4

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// fatal error all groutines are asleep - deadlock!
// that is because WaitGroup run wait more one
func Waitgroup() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 10; i < 10; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&count))
}
