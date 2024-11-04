package python

// #cgo pkg-config: python3-embed
// #include <Python.h>
import "C"
import "unsafe"

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
