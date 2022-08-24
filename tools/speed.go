package tools

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func doWhat() {
	for i := 0; i <= 1000; i++ {
		j := i % 5
		fmt.Println("get mold value", j)
	}
}

func SpeedTime(handler func()) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	// get function name by reflection
	funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	fmt.Println(funcName+"spend time:", elapsed)
}
