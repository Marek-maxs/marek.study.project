package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gofrs/uuid"
)
/**
*
* Author: Marek
* Date: 2023-03-09 14:18
* Email: 364021318@qq.com
*
 */

type TE struct {
	Name string
}
func main() {
	str := uuid.FromStringOrNil("bc33674b-0cac-41a3-9d69-16d121619343")
	encodeString := base64.StdEncoding.EncodeToString(str.Bytes())
	fmt.Println(encodeString)
	arr := make(map[string]map[string]TE)
	//arr := map[string]map[string]string{"aa":{"bb":"333", "cc":"555"}}
	//arr["aa"]["ff"] = "99"
	if arr["a"] == nil {
		arr["a"] = make(map[string]TE)
	}
	newT := TE{Name:"11"}
	arr["a"]["b"] = newT
	newV := TE{Name:"22"}
	arr["a"]["v"] = newV
	Rt(arr)
	fmt.Println(arr)
}

func Rt(varr map[string]map[string]TE) {

	//arr := map[string]map[string]string{"aa":{"bb":"333", "cc":"555"}}
	//arr["aa"]["ff"] = "99"
	if varr["a"] == nil {
		varr["a"] = make(map[string]TE)
	}
	newT := TE{Name:"11"}
	varr["a"]["g"] = newT
	newV := TE{Name:"22"}
	varr["a"]["h"] = newV
}