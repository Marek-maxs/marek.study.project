package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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

// read csv file context and print

func main() {
	// 使用 csv.NewReader 来解析 CSV 数据
	file, err := os.Open("D:\\company\\EN-TRAK-company\\hub-server\\starbucks-HK\\stb-need-upload-raw-data.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(records) == 0 {
		fmt.Println("empty row")
		return
	}

	datapointNameMap := make(map[int]string)
	for k, col := range records[0] {
		if k >= 2 {
			datapointNameMap[k] = col
		}
	}

	for _, record := range records {
		if record[0] == "timestamp" {
			continue
		}

		var timeCol int64
		var strTime string
		// 处理每一行，即使列数不固定
		for i, field := range record {
			if i == 0 {
				strTime = field

				parsedTime, err := time.ParseInLocation("2006-01-02T15:04", field, locationFromOffset(480))
				if err != nil {
					fmt.Println("parse datapoint time failed:", err.Error())

					continue
				}

				timeCol = parsedTime.Unix()
				continue
			}

			if i >= 2 {
				datapointName, ok := datapointNameMap[i]
				if !ok {
					continue
				}

				value, err := strconv.ParseFloat(field, 64)
				if err != nil {
					fmt.Println("parser coll value failed:", err.Error())

					value = 0
				}

				//data := fmt.Sprintf(dataFormat, timeCol, serial, strings.TrimSpace(datapointName), value*scale)
				//
				//body, err := sjson.SetBytes([]byte("{}"), "Data", data)
				//if err != nil {
				//	fmt.Printf("data: %s set data error: %v\n", data, err)
				//	continue
				//}

				fmt.Printf("strTime:%s,data wrote to %s: %f\n",
					strTime, strings.TrimSpace(datapointName), value*1000)
				fmt.Println(timeCol)
			}
		}
	}

	fmt.Println("End run lambda program")
}
