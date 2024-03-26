// nolint
package python

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

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

func GetExpressionValue(expression string) string {
	Py_Initialize()
	defer Py_Finalize()
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(dir)
	// 3. 导入expression模块
	hello := ImportModule("./python/file", "expression")
	defer hello.DecRef() // pyObject => string 解析结果
	_, err = pythonRepr(hello)
	if err != nil {
		log.Error().Err(err).Str("expression", expression).
			Msg("pythonRepr failed")

		return ""
	}
	// 5. 获取函数方法
	GetNodeValue := hello.GetAttrString("GetNodeValue")
	defer GetNodeValue.DecRef()
	// 设置调用的参数（一个元组）
	nargs := PyTuple_New(1)
	// 创建存储空间
	defer nargs.DecRef()
	input := PyUnicode_FromString(expression)
	// input不需要DecRef，因为DecRef args的时候就一起DecRef了
	PyTuple_SetItem(nargs, 0, input)
	// 设置值
	res := GetNodeValue.Call(nargs, Py_None)
	// 调用
	result := PyUnicode_AsUTF8(res)

	log.Info().Str("expression", expression).Str("res", result).
		Msg("python calculate")

	return result
}
