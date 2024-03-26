package main

// #cgo pkg-config: python3
// #include <Python.h>
import "C"
import (
   "fmt"
   "unsafe"
)

type CChar C.char

func (p *CChar) GoString() string {
   return C.GoString((*C.char)(p))
}

func PrintCString(cs *C.char) {
   ccc := (*CChar)(cs)
   print(ccc.GoString(), "\n")
}

func PrintGoString(cs string) {
   var aa = []byte(cs)
   ccc := (*C.char)(unsafe.Pointer(&aa[0]))
   PrintCString(ccc)

   bb := C.CString(cs)
   PrintCString(bb)
}

func ab() {

   pycodeGo := `
 import sys
 for path in sys.path:
   print(path)
`

   defer C.Py_Finalize()
   C.Py_Initialize()
   pycodeC := C.CString(pycodeGo)
   defer C.free(unsafe.Pointer(pycodeC))
   res := C.PyRun_SimpleString(pycodeC)
   type PyObject C.PyObject
   fmt.Println("goalng:",res)
   data :=  (*C.char)(unsafe.Pointer(res))
C.GoString(pycodeC)
   fmt.Println("goalng pycodeC:",data.GoString())
   fmt.Println(C.GoString(pycodeC))
}

func PyRun_SimpleString(command string) int {
   commandC := C.CString(command)
   defer C.free(unsafe.Pointer(commandC))
   return int(C.PyRun_SimpleString(commandC))
}