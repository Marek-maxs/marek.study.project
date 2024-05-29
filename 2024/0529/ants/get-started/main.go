package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}

	t.wg.Done()
}

func taskFunc(data interface{}) {
	task := data.(*Task)
	task.Do()

	fmt.Printf("task:%d sum: %d\n", task.index, task.sum)
}

const (
	DataSize    = 10000
	DataPerTask = 100
)

func main() {
	fmt.Println("start run program")
	// 创建一个大小为10的线程池，用来执行指定的一种function
	p, _ := ants.NewPoolWithFunc(10, taskFunc)
	defer p.Release()

	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	tasks := make([]*Task, 0, DataSize/DataPerTask)

	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{
			index: i + 1,
			nums:  nums[i*DataPerTask : (i+1)*DataPerTask],
			sum:   0,
			wg:    &wg,
		}

		tasks = append(tasks, task)

		p.Invoke(task) // invoke submit a task to pool(将任务提交到池)
	}

	wg.Wait()
	fmt.Printf("running groutines: %d\n", ants.Running())

	var sum int
	for _, task := range tasks {
		sum += task.sum
	}

	var expect int
	for _, num := range nums {
		expect += num
	}

	fmt.Printf("finish all tasks, result is %d\n", sum, expect)
}
