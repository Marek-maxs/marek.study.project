package main

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.Py_Initialize()

	defer C.Py_Finalize()

	// 执行 Python 代码
	code := C.CString(`
print('Hello, Python!')
`)

	defer C.free(unsafe.Pointer(code))

	C.PyRun_SimpleString(code)

	// 调用 Python 模块
	moduleName := C.CString("os")

	defer C.free(unsafe.Pointer(moduleName))

	module := C.PyImport_ImportModule(moduleName)

	if module == nil {
		fmt.Println("Failed to import module")
		return
	}

	// 调用 Python 模块中的函数
	functionName := C.CString("getcwd")

	defer C.free(unsafe.Pointer(functionName))

	function := C.PyObject_GetAttrString(module, functionName)

	if function == nil {
		fmt.Println("Failed to get function")
		return
	}

	args := C.PyTuple_New(0)

	result := C.PyObject_CallObject(function, args)

	if result == nil {
		fmt.Println("Failed to call function")
		return
	}

	defer C.Py_DecRef(result)

	value := C.PyUnicode_AsUTF8(result)

	fmt.Println(value)
}