package main

/**
*
* Author: Marek
* Date: 2024-02-15 16:53
* Email: 364021318@qq.com
*
 */

type Job struct {
	Id int
	// 需要计算的随机数
	RandNum int
}

func NewJob(id , randNum int) *Job {
	return &Job{
		Id:      id,
		RandNum: randNum,
	}
}

func (j *Job) GetSum() int {
	// 随机数接过来
	randNum := j.RandNum
	// 随机数每一位相加
	// 定义返回值
	var sum int

	for randNum != 0 {
		tmp := randNum % 10
		sum += tmp
		randNum /= 10
	}

	return sum
}