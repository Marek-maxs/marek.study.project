package main

import "fmt"

// 通过巧妙地使用空接口，闭包和高阶函数，我们能实现一个通用的惰性生产器工厂函数

type Any interface {

}

type Evalfunc func(Any)(Any, Any)

// 相当于主函数
func GetLazyintevaluator() {
	// 计算偶数的函数 接收一个初始值 即从它开始计算偶数
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	// 生成一个新的int 生成器
	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Println("%vth even: %v \n", i, even())
	}
}

func BuildLazyIntEvaluator(evalFunc Evalfunc, initState Any) func()int  {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}

}

func BuildLazyEvaluator(evalFunc Evalfunc, initState Any) func()Any  {
	// 返回值通道
	retValChan := make(chan Any)
	// 循环函数， 这个函数是生成值的
	loopFunc := func() {
		// 初始状态
		var actState = initState
		// 返回值
		var retVal Any
		// 循环
		for {
			// 经过实时计算的返回值和实际状态
			retVal, actState = evalFunc(actState)
			// 返回值写入通道
			retValChan <- retVal
		}
	}
	// 返回值函数， 该函数专门用于读取retValChan 通道的值
	// 这个函数是读取值的
	retFunc := func() Any {
		return <- retValChan
	}
	// 协程方式启动循环函数，该函数用于往通道写值
	go loopFunc()
	return retFunc
}