package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func wrapper(i int, wg *sync.WaitGroup) func() {
	return func() {
		fmt.Printf("hello from task: %d\n", i)

		wg.Done()
	}
}

func panicHandler(err interface{}) {
	fmt.Println(os.Stderr, err)
}

func main() {
	p, _ := ants.NewPool(2, ants.WithPanicHandler(panicHandler))
	defer p.Release()

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		err := p.Submit(wrapper(i, &wg))
		if err != nil {
			fmt.Println("pool submit err:", err.Error())
		}

		fmt.Printf("submit input value: %d\n", i)
	}

	time.Sleep(1 * time.Second)
	err := p.Submit(wrapper(3, &wg))
	if err != nil {
		fmt.Println("pool submit err:", err.Error())
	}

	wg.Wait()
}
