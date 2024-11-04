package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/**
*
* Author: Marek
* Date: 2023-02-17 1:13
* Email: 364021318@qq.com
*
 */
const AddActTimeLayout = "2006-01-02 15:04"
type AddActiveDateTime struct {
	time.Time
}

//自定义反序列化方法，实现UnmarshalJSON()接口
func (dt *AddActiveDateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")             //去掉首尾的"
	dt.Time, err = time.Parse(AddActTimeLayout, s) //格式化时间格式
	// fmt.Printf("[dt.Time:]=%v\n", dt.Time)
	return
}

type AddActive struct {
	Title        string            `json:"title"`
	StartTime    AddActiveDateTime `json:"start_time"`
}

type Txter interface {
	TxtJSON([]byte) error
}

func main() {
	var activeInfo AddActive
	body := `{"title": "2323", "start_time": "2023-02-17"}`
	err := json.Unmarshal([]byte(body), &activeInfo)
	fmt.Println(err)
}

func Txt() {

}