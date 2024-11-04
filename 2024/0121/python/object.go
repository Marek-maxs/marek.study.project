package python

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"
import "unsafe"

//None : https://docs.python.org/3/c-api/none.html#c.Py_None
var Py_None = togo(C.Py_None)

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
