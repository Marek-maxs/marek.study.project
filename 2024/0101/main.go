package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/spf13/cast"
	"time"
)

/**
*
* Author: Marek
* Date: 2024-01-01 22:11
* Email: 364021318@qq.com
*
 */

type One struct {
	A string
	B string
}

func main() {
	// get 4 value
	data := []byte{0x12, 0x34, 0x56, 0x78}
	fmt.Printf("%b\n", data)

	for i, b := range data {
		data[i] = (b&0x0F)<<4 | (b&0xF0)>>4
	}

	fmt.Println(data)

	povID := uuid.FromStringOrNil("d777481e-8594-4ef9-ae8e-2313a09c2301")
	newPov := base64.StdEncoding.EncodeToString(povID.Bytes())
	fmt.Println(newPov)

	pasePov, _ := base64.StdEncoding.DecodeString("5vtqmXY/Rv6mh/svHFIa+w==")
	fmt.Println(uuid.FromBytesOrNil(pasePov).String())

	// 获取当前时间的年、月、日信息
	now := time.Now()

	// 如果今天不是周一或者超过了本周的最后一天（周六），则向前移动直至遇到周一
	//for i := 0; i < int(now.Weekday()); i++ {
	//	now = now.AddDate(0, 0, -i-1)
	//	fmt.Println(now.String())
	//}
	now = now.AddDate(0, 0, -int(now.Weekday())+7)
	fmt.Println(now.String())

	ts := 1704124800
	a := time.Unix(int64(ts), 0)
	loc,_ := time.LoadLocation("Asia/Shanghai")

	fmt.Println(a.In(loc).String())
fmt.Println("-------")
	datas := []byte{0x00, 0x00, 0xc9, 0x20}

	fmt.Println(cast.ToFloat32(binary.BigEndian.Uint32(datas)))
	aa := 1706796144
	b := 1706758444

	cc := 1706796144

	if cc > b {
		cc = b
	}
	if cc > aa {
		cc = aa
	}

	fmt.Println("cc", cc)

	arr := make(map[int]*One)

	if arr[5] == nil {
		arr[5] = new(One)
	}
	arr[5].A = "222"
	fmt.Println(arr[5])

}