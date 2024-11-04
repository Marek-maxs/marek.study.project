package main

import (
	"fmt"
	"regexp"
	"time"
)

// isValidString 检查字符串是否只包含数字和特定的符号（+，-，*，%，/，.）
func isValidString(s string) bool {
	// 正则表达式，匹配数字和符号
	re := regexp.MustCompile(`^[\d\+\-\*\%\/\.]*$`)
	return re.MatchString(s)
}

func main() {
	//testStrings := []string{
	//	//	"{1088}+{1088}-10*{323}/2+{3232}%{323}",
	//	//	"hello123",
	//	//	"1088a",
	//	//	"1088-1088*323%323/1",
	//	//	"{1088}+10*{323}-2+{3232}%{323}",
	//	//	"1088+1088*3.23-2/1", // 允许小数点，如果需要的话
	//	//}
	//	//
	//	//for _, str := range testStrings {
	//	//	if isValidString(str) {
	//	//		fmt.Printf("字符串 '%s' 是有效的\n", str)
	//	//	} else {
	//	//		fmt.Printf("字符串 '%s' 是无效的\n", str)
	//	//	}
	//	//}

	GenerateTime(120)
}

func GenerateTime(minutes int) {
	loc,_ := time.LoadLocation("Asia/Hong_Kong")
	curTime := time.Now().Add(time.Minute * time.Duration(-minutes))
	totalHours := minutes / 60
	startTimestamp := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), curTime.Hour(),
		0, 0, 0, loc)

	for i := 0; i < totalHours; i++ {
		nextHour := time.Hour * time.Duration(i)
		nextPoint := startTimestamp.Add(nextHour)
		nextPointAddMinutes := nextPoint.Add(59 * time.Minute)

		fmt.Println(nextPoint.Format(time.DateTime))
		fmt.Println(nextPointAddMinutes.Format(time.DateTime))
	}
}