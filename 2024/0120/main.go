package main
import (
	"fmt"
	"github.com/datadog/go-python3"
	"os"
	)
func init() {
	// 1. 初始化python环境
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}
}
func main() {
	// 7. 结束环境(提前defer)
	defer python3.Py_Finalize()
	// 2. 设置python import 的路径
	p := ""
	InsertBeforeSysPath(p)
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
	 args := python3.PyTuple_New(1)
	 // 创建存储空间
	 defer args.DecRef()
	 input := python3.PyUnicode_FromString("xwj")
	 // input不需要DecRef，因为DecRef args的时候就一起DecRef了
	 python3.PyTuple_SetItem(args, 0, input)
	 // 设置值
	 res := SayHello.Call(args, python3.Py_None)
	 // 调用
	 fmt.Printf("[FUNC] res = %s\n", python3.PyUnicode_AsUTF8(res))
	 // 6. 调用第三方库 sklearn
	 sklearn := hello.GetAttrString("sklearn")
	 defer sklearn.DecRef()
	 skVersion := sklearn.GetAttrString("__version__")
	 defer skVersion.DecRef()
	 sklearnRepr, err := pythonRepr(sklearn)
	 if err != nil {
	 	panic(err)
	 }
	 skVersionRepr, err := pythonRepr(skVersion)
	 if err != nil {
	 	panic(err)
	 }
	 fmt.Printf("[IMPORT] sklearn = %s\n", sklearnRepr)
	 fmt.Printf("[IMPORT] sklearn version = %s\n", skVersionRepr)
}
// InsertBeforeSysPath
// @Description: 添加site-packages路径即包的查找路径
// @param p
func InsertBeforeSysPath(p string){
	sysModule := python3.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	pObject := python3.PyUnicode_FromString(p)
	defer pObject.DecRef()
	python3.PyList_Append(path, pObject)
}
// ImportModule
// @Description: 倒入一个包
// @param dir
// @param name
// @return *python3.PyObject
func ImportModule(dir, name string) *python3.PyObject {
	sysModule := python3.PyImport_ImportModule("sys")
	// import sys
	path := sysModule.GetAttrString("path")
	// path = sys.path
	dirObject := python3.PyUnicode_FromString(dir)
	defer dirObject.DecRef()
	python3.PyList_Insert(path, 0, dirObject)
	// path.insert(0, dir)
	return python3.PyImport_ImportModule(name) // return __import__(name)
}
// pythonRepr
// @Description: PyObject转换为string
// @param o
// @return string
// @return error
func pythonRepr(o *python3.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil")
	}
	s := o.Repr() // 获取对象转换为可读
	 if s == nil {
	 	python3.PyErr_Clear()
	 	return "", fmt.Errorf("failed to call Repr object method")
	 }
	 defer s.DecRef()
	return python3.PyUnicode_AsUTF8(s), nil
}
// PrintList
// @Description: 输出一个List
// @param list
// @return error
func PrintList(list *python3.PyObject) error {
	if exc := python3.PyErr_Occurred();
	list == nil && exc != nil {
		return fmt.Errorf("Fail to create python list object")
	}
	defer list.DecRef()
	repr, err := pythonRepr(list)
	if err != nil {
		return fmt.Errorf("fail to get representation of object list")
	}
	fmt.Printf("python list: %s\n", repr)
	return nil
}