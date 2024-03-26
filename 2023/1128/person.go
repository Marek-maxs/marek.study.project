package main

import (
	"encoding/json"
	"time"
)

type People struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Birthday Time `json:"_"`
}

func (p *People) UnmarshalJSON(b []byte) error {
	// 字义临时类型， 用来接受非 `json:"_"`的字段
	type tmp People
	// 用中间变量接收json串， tmp 以外的字段用来接受 `json:"_"` 属性字段
	var s = &struct {
		tmp
		// string 先接收字符串类型， 一会转换
		Birthday string `json:"birthday"`
	}{}
	// 解析
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	localTimeObj, err := time.ParseInLocation(timeFormat, s.Birthday, time.Local)
	if err != nil {
		return err
	}

	s.tmp.Birthday = localTimeObj
	// tmp 类型转换回 People， 并赋值
	*p = People(s.tmp)

	return nil
}