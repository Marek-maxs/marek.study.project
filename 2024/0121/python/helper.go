package python

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *PyObject {
	return (*PyObject)(cobject)
}

func toc(object *PyObject) *C.PyObject {
	return (*C.PyObject)(object)
}
