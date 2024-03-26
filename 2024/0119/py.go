package main

/*
#cgo CFLAGS: -I/usr/include/python3.10
#cgo LDFLAGS: -lpython3.10
#include <Python.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// 调用Python解释器执行Python代码并获取执行结果
	code := []byte(`print("Hello from Python!")`)
	cCode := C.CString(string(code))
	defer C.free(unsafe.Pointer(cCode))

	cResult := C.PyRun_SimpleString(cCode)
	if cResult != nil {
		// 将Python结果转换为Go字符串并打印出来
		result := C.GoString(cResult)
		fmt.Println("执行结果:", result)
	} else {
		fmt.Println("Python代码执行出错.")
	}
}