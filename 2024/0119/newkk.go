package main

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"

import (
	"fmt"
	"unsafe"
)

//PyObject : https://docs.python.org/3/c-api/structures.html?highlight=pyobject#c.PyObject
type PyObject C.PyObject

//GetAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttrString
func (pyObject *PyObject) GetAttrString(attr_name string) *PyObject {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return togo(C.PyObject_GetAttrString(toc(pyObject), cattr_name))
}

//DecRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_DECREF
func (pyObject *PyObject) DecRef() {
	C.Py_DecRef(toc(pyObject))
}

//Repr : https://docs.python.org/3/c-api/object.html#c.PyObject_Repr
func (pyObject *PyObject) Repr() *PyObject {
	return togo(C.PyObject_Repr(toc(pyObject)))
}

//Call : https://docs.python.org/3/c-api/object.html#c.PyObject_Call
func (pyObject *PyObject) Call(args *PyObject, kwargs *PyObject) *PyObject {
	return togo(C.PyObject_Call(toc(pyObject), toc(args), toc(kwargs)))
}

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *PyObject {
	return (*PyObject)(cobject)
}

func toc(object *PyObject) *C.PyObject {
	return (*C.PyObject)(object)
}

//None : https://docs.python.org/3/c-api/none.html#c.Py_None
var Py_None = togo(C.Py_None)

//PyImport_ImportModule : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModule
func PyImport_ImportModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_ImportModule(cname))
}

//PyUnicode_FromString : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_FromString
func PyUnicode_FromString(u string) *PyObject {
	cu := C.CString(u)
	defer C.free(unsafe.Pointer(cu))

	return togo(C.PyUnicode_FromString(cu))
}

//PyList_Insert : https://docs.python.org/3/c-api/list.html#c.PyList_Insert
func PyList_Insert(p *PyObject, index int, item *PyObject) int {
	return int(C.PyList_Insert(toc(p), C.Py_ssize_t(index), toc(item)))
}

//PyErr_Clear : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Clear
func PyErr_Clear() {
	C.PyErr_Clear()
}

//PyUnicode_AsUTF8 : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_AsUTF8
func PyUnicode_AsUTF8(unicode *PyObject) string {
	cutf8 := C.PyUnicode_AsUTF8(toc(unicode))
	return C.GoString(cutf8)
}

//PyTuple_New : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_New
func PyTuple_New(len int) *PyObject {
	return togo(C.PyTuple_New(C.Py_ssize_t(len)))
}

//PyTuple_SetItem : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_SetItem
func PyTuple_SetItem(p *PyObject, pos int, o *PyObject) int {
	return int(C.PyTuple_SetItem(toc(p), C.Py_ssize_t(pos), toc(o)))
}

// ImportModule
// @Description: 倒入一个包
// @param dir
// @param name
// @return *python3.PyObject
func ImportModule(dir, name string) *PyObject {
	sysModule := PyImport_ImportModule("sys")
	// import sys
	path := sysModule.GetAttrString("path")
	// path = sys.path
	dirObject := PyUnicode_FromString(dir)
	defer dirObject.DecRef()
	PyList_Insert(path, 0, dirObject)
	// path.insert(0, dir)
	return PyImport_ImportModule(name) // return __import__(name)
}

// pythonRepr
// @Description: PyObject转换为string
// @param o
// @return string
// @return error
func pythonRepr(o *PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil")
	}
	s := o.Repr() // 获取对象转换为可读
	if s == nil {
		PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method")
	}
	defer s.DecRef()
	return PyUnicode_AsUTF8(s), nil
}

func main() {
	C.Py_Initialize()

	defer C.Py_Finalize()

	// 执行 Python 代码
	code := C.CString(`
print('Hello, Python!')
`)

	defer C.free(unsafe.Pointer(code))

	C.PyRun_SimpleString(code)

	// 3. 导入hello模块
	hello := ImportModule("./hello", "hello")
	defer hello.DecRef() // pyObject => string 解析结果
	helloRepr, err := pythonRepr(hello)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[MODULE] repr(hello) = %s\n", helloRepr)
	// 4. 获取变量
	a := hello.GetAttrString("a")
	defer a.DecRef()
	aString, err := pythonRepr(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[VARS] a = %#v\n", aString)
	// 5. 获取函数方法
	SayHello := hello.GetAttrString("SayHello")
	defer SayHello.DecRef()
	// 设置调用的参数（一个元组）
	args := PyTuple_New(1)
	// 创建存储空间
	defer args.DecRef()
	input := PyUnicode_FromString("xwj")
	// input不需要DecRef，因为DecRef args的时候就一起DecRef了
	PyTuple_SetItem(args, 0, input)
	// 设置值
	res := SayHello.Call(args, Py_None)
	// 调用
	fmt.Printf("[FUNC] res = %s\n", PyUnicode_AsUTF8(res))
}