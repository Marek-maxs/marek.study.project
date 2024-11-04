package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
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
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Invalid arguments")
		return
	}
	// get 4 value
	data := []byte{0x12, 0x34, 0x56, 0x78}
	//fmt.Printf("%b\n", data)

	for i, b := range data {
		data[i] = (b&0x0F)<<4 | (b&0xF0)>>4
	}

	fmt.Println("get arg input:", args[0])
	povID := uuid.FromStringOrNil(args[0])
	newPov := base64.StdEncoding.EncodeToString(povID.Bytes())
	fmt.Println("base64 ID:", newPov)

	//fmt.Println(time.Unix(1714492800, 0).UTC().Format(time.DateTime))
	//fmt.Println(time.Unix(1717171199, 0).UTC().Format(time.DateTime))
	//
	pasePov, _ := base64.StdEncoding.DecodeString(newPov)
	fmt.Println("restore uuid:", uuid.FromBytesOrNil(pasePov).String())
	//
	//value := float32(18.48564)
	//eui := float32(0)
	//fmt.Println(value / eui)
	//// 获取当前时间的年、月、日信息
	//now := time.Now()
	//
	//// 如果今天不是周一或者超过了本周的最后一天（周六），则向前移动直至遇到周一
	////for i := 0; i < int(now.Weekday()); i++ {
	////	now = now.AddDate(0, 0, -i-1)
	////	fmt.Println(now.String())
	////}
	//now = now.AddDate(0, 0, -int(now.Weekday())+7)
	//fmt.Println(now.String())
	//
	//ts := 1704124800
	//a := time.Unix(int64(ts), 0)
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//
	//fmt.Println(a.In(loc).String())
	//fmt.Println("-------")
	//fmt.Println(loc.String())
	//newLoc, _ := locationFromOffset(480)
	//fmt.Println(newLoc.String())
	//datas := []byte{0x00, 0x00, 0xc9, 0x20}
	//
	//fmt.Println(cast.ToFloat32(binary.BigEndian.Uint32(datas)))
	//aa := 1706796144
	//b := 1706758444
	//
	//cc := 1706796144
	//
	//if cc > b {
	//	cc = b
	//}
	//if cc > aa {
	//	cc = aa
	//}
	//
	//fmt.Println("cc", cc)
	//
	//arr := make(map[int]*One)
	//
	//if arr[5] == nil {
	//	arr[5] = new(One)
	//}
	//arr[5].A = "222"
	//fmt.Println(arr[5])

}

func locationFromOffset(offset int32) (*time.Location, int32) {
	offsetHour := offset / 60

	locName := ""
	if offsetHour >= 0 {
		locName = fmt.Sprintf("UTC+%d", offsetHour)
	} else { // minus
		locName = fmt.Sprintf("UTC%d", offsetHour)
	}

	return time.FixedZone(locName, int(offset)*60), offsetHour
}
