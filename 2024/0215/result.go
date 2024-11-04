package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-02-15 16:57
* Email: 364021318@qq.com
*
 */

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func NewResult(job *Job, sum int) *Result {
	return &Result{
		job: job,
		sum: sum,
	}
}

func (r *Result) PrintLog() {
	fmt.Printf("job id:%v randnum: %v result: %d\n", r.job.Id,
		r.job.RandNum, r.sum)
}