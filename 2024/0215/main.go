package main

import (
	"math/rand"
)

/**
*
* Author: Marek
* Date: 2024-02-15 16:06
* Email: 364021318@qq.com
*
 */

/*
Goroutine 池

worker pool

- 本质上是生产者模型
- 可以有效的控制 goroutine 数量， 防止暴涨
- 需求：
	计算一个数字的各个位数之和，例如数字123， 结果 1+2+3 = 6
	随机生成数字进行计算
*/

func main() {
	// 需要2个管道
	// 1.job管理
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(64, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			result.PrintLog()
		}
	}(resultChan)

	var id int
	// 循环创建job, 输入到管道
	for {
		id++
		// 生成随机数
		randIntNum := rand.Int()
		job := NewJob(id, randIntNum)

		jobChan <- job
		// 跳出循环
		if id > 1000 {
			break
		}
	}
}

// 创建工作池
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				sum := job.GetSum()

				// 想要的结果是Result
				r := NewResult(job, sum)

				// 运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}