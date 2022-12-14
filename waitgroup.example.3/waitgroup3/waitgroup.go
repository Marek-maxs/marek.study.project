package waitgroup3

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// negative WaitGroup counter
// that is because WaitGroup Done run more one
func Waitgroup() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}
	wg.Done()
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&count))
}
