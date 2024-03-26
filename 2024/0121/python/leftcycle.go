package python

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"

//Py_Initialize : https://docs.python.org/3/c-api/init.html#c.Py_Initialize
func Py_Initialize() {
	C.Py_Initialize()
}

//Py_Finalize : https://docs.python.org/3/c-api/init.html#c.Py_Finalize
func Py_Finalize() {
	C.Py_Finalize()
}
