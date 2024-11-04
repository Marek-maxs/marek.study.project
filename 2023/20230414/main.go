package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/tidwall/sjson"
	"strconv"
	"strings"
	"time"
)

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

type LiveDataTime struct {
	M00 float64 `json:"m00,omitempty"`
	M01 float64 `json:"m01,omitempty"`
	M02 float64 `json:"m02,omitempty"`
	M03 float64 `json:"m03,omitempty"`
}

func main() {
	maxGoroutinesNum := 5
	ticketPool := make(chan int, maxGoroutinesNum)

	for i := 0; i < maxGoroutinesNum; i++ {
		ticketPool <- i
	}

	// 这种写法会出现 deadlock 现像
	// for conn := range ticketPool {
	// 	fmt.Println(conn)
	// }

	// 这种写法才是正常的
	for i := 0; i < maxGoroutinesNum; i++ {
		conn := <-ticketPool
		fmt.Println(conn)
	}
	usages := `{"2023-06-02 03:00:00 +0000 UTC":{"m00":0.138,"m01":0.138,"m02":0.154,"m03":0.124}}`
	var a map[string]map[string]float64
	json.Unmarshal([]byte(usages), &a)

	fmt.Println(a)
	fmt.Println("---------------")
	// use base64 encode uuid
	povID := uuid.FromStringOrNil( "05398736-e33c-4d6e-83c2-720eb0793718")
	newPovID := base64.StdEncoding.EncodeToString(povID.Bytes())
	//decoded, err := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(newPovID))
	m := make(map[string]interface{})
	m["list"] = []map[string]int{
		{"cid": 6, "aid": 0, "value": 1},
	}
	s, _ := sjson.Set(addSceneValueJSON, "argument", m)
	fmt.Println(s)

	//nowTime := time.Now().Unix()
	//nowDate := time.Unix(nowTime, 0)
	fmt.Println(0%5)
	nowDate := time.Unix(1682402835, 0)

	difference := time.Since(nowDate)
	fmt.Println(difference.Minutes())

	// 创建channel
	ch := make(chan string, 10)

	l := len(ch)
	fmt.Println("1通道的长度：", l)
	if l > 0 {
		for i := 0; i < l; i++ {
			fmt.Println(<-ch)
		}
	}

	// 发送数据
	for i := 0; i<5;i++ {
		ch <- fmt.Sprintf("data%d", i)
	}

	l = len(ch)
	fmt.Println("2通道的长度：", l)
	// 接收数据
	for i := 0; i < l; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(time.Now().Format("200601021504"))

	n := "-2"
	newN, _ := strconv.Atoi(n)
	fmt.Println(newN)
	fmt.Println(fmt.Sprintf("%.2f", float32(10)/100))
	boyCost, _ := strconv.ParseFloat("0.1", 64)
	fmt.Println(boyCost)
	//fmt.Println("fwfwe:",int64(float32(10)/100))

	fmt.Println(0.912+0.6559999999999999+1.168+1.164+1.16+1.416+1.156+ 1.16+ 1.064+1.3199999999999998+1.168+0.912+0.912+ 1.164+1.1640000000000001+1.164+ 1.1599999999999997+0.8679999999999999+ 0.5880000000000001+0.8599999999999999+0.604+ 0.916+0.6599999999999999+0.912+ 0.92+1.1720000000000002+0.9119999999999998+1.1760000000000002+1.172+1.168+1.172+1.1719999999999997+0.9119999999999999+ 1.176+0.912+1.172+1.248+0.9199999999999999+0.9119999999999999+ 0.888+0.968+ 1.016+0.916+0.9119999999999999+0.9119999999999999+1.172+0.788+1.22+1.212+ 1.212+ 1.1119999999999999+1.156+1.412+1.152+1.156+ 1.156+ 1.1560000000000001+1.1560000000000001+0.9+1.1520000000000001)

	fmt.Println(0.912+0.656+1.168+1.1639999999999997+1.16+1.416+1.1559999999999997+1.1599999999999997+1.0639999999999998+1.3199999999999998+1.168+0.9120000000000001+0.912+1.164+1.1640000000000001+1.1639999999999997+1.1599999999999997+0.8679999999999999+0.588+0.86+0.6040000000000001+0.916+0.66+0.9120000000000001+0.92+1.1720000000000002+0.912+1.176+1.1719999999999997+1.1679999999999997+1.172+1.172+0.912+1.176+0.912+1.1719999999999997+1.248+0.9199999999999999+0.912+0.888+0.9680000000000001+1.016+0.916+0.912+0.912+1.1720000000000002+0.7880000000000001+1.2200000000000002+1.212+1.2119999999999997+1.1119999999999999+1.156+1.412+1.1519999999999997+1.156+1.156+1.1560000000000001+1.1560000000000001+0.9+1.1520000000000001)

	fmt.Println(int(time.Unix(1688655600, 0).Weekday()))

	fmt.Println(fmt.Sprintf("%02d-%02d-%02d", 2023, 7, 7))

	ts := 1687651200
	loc, _ := time.LoadLocation("Asia/Hong_Kong")
	datetime := time.Unix(int64(ts), 0 ).In(loc)
	weekday := datetime.Weekday()
	fmt.Println(int(weekday))
	fmt.Println(weekday)
	fmt.Println(CheckIsWorkDay(int(weekday)))

	str := "/gw/acrelHW/AWT100/data/12308017890017"
	newStr := strings.SplitN(str, "/", -1)
	fmt.Println(str[:len(str)-len(newStr[len(newStr)-1])])

	kwh := 345.5654375
	if kwh < 100000 && -100000 < kwh {
		fmt.Println(kwh, "sdfsfwefwef")
	}
	fmt.Println("end")

	data := time.Now().In(loc).Minute()
	fmt.Println(data)
	fmt.Println(uuid.Must(uuid.NewV4()))
}

func CheckIsWorkDay(week int) bool {
	switch week {
	case 6:
		return false
	case 0:
		return false
	}

	return true
}

const addSceneValueJSON = `{
  "type": "scene",
  "command": "add",
  "argument": {
	"ieee":"C2A08F01006F0D00",
	"ep":1,
	"gid": 1,
	"sid": 1,
	"name":"name1",
	"list":[{
	"cid":6,
	"aid":0,
	"value":1
	}]
  },
  "session": "afkcdbb6xsnk7ra",
  "sequence": 1017
}`
