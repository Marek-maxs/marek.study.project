package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main () {
	msg := "广州天河体育场羽毛球官1\\\\r\\\\n开始时间:2023-06-24 12:00\\\\r\\\\n结束时间:2023-06-24 13:00\\\\r\\\\n组织者:Marek\\\\r\\\\n\\\\r\\\\n已报名:[2男1女][召集中]\\\\r\\\\n\\\\r\\\\n1.Marek3\\r\\n2.Marek2(女)\\r\\n3.Marek1\\r\\n&club=9"

	data := url.Values{}
	data.Set("name", msg)
	_, err := http.PostForm("http://127.0.0.1:8889/send/message", data)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Unix(1687013880, 0).Format("2006-01-02 15:04"))

	i := 0

	for {
		i++
		if i == 10 {
			fmt.Println("break for log")
			break
		}
	}
}