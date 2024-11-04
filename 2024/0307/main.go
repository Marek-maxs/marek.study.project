package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 获取本周的开始时间（周一）
	startOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -(int(now.Weekday()) - 1))

	// 获取本周的结束时间（周日）
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	// 打印结果
	fmt.Println("本周的开始时间（周一）:", startOfWeek)
	fmt.Println("本周的结束时间（周日）:", endOfWeek)

	//loc,_ := time.LoadLocation("Asia/Shanghai")
	loc := locationFromOffset(480)
	fmt.Println(loc.String())
}

func locationFromOffset(offset int32) *time.Location {
offsetHour := offset / 60

locName := ""
if offsetHour >= 0 {
locName = fmt.Sprintf("UTC+%d", offsetHour)
} else { // minus
locName = fmt.Sprintf("UTC%d", offsetHour)
}

return time.FixedZone(locName, int(offset)*60)
}