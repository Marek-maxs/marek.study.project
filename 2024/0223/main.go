package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const path  = `E:\Tool\data1.csv`
const path2  = `E:\Tool\data2.xlsx`

const laytime_local  = "2006-01-02T15:04"

func parserCsv() {
	// 打开CSV文件
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建新的CSV阅读器
	reader := csv.NewReader(file)

	record, err := reader.Read()
	if err != nil {
		log.Fatal("can not reader csv file info")
	}

	datapointNameMap := make(map[int]string)
	for k, col := range record {
		if k >= 2 {
			datapointNameMap[k] = col
			fmt.Println(col)
		}
	}

	var parserMsg []Msg
	// 读取文件的每一行
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if record[0] == "timestamp" {
			continue
		}

		var timeCol int64
		// 处理每一行，即使列数不固定
		for i, field := range record {
			if i == 0 {
				parsedTime, err := time.Parse(laytime_local, field)
				if err != nil {
					fmt.Println("parse datapoint time failed:",err.Error())
					return
				}

				timeCol = parsedTime.Unix()
				continue
			}
			if i >= 2 {
				value, err := strconv.ParseFloat(field, 64)
				if err != nil {
					fmt.Println("parser coll value failed:", err.Error())

					value = 0
				}
				parserMsg = append(parserMsg, Msg{
					Time:          timeCol,
					DatapointName: datapointNameMap[i],
					Value:         value,
				})
			}
		}
	}

	for _, val := range parserMsg {
		fmt.Println("time:", val.Time, " DatapointName:", val.DatapointName, " Value:", val.Value)
	}
}


type Msg struct {
	Time int64
	DatapointName string
	Value float64
}

func main() {
	// excel()
	//parserCsv()
	_, filename, ok := strings.Cut("meters_data/starbucks-uploads-meter-01_0_HcBUDBpTm.csv", "/")
	if !ok {
		fmt.Println("not filename")
	}
	arr := strings.Split(filename, "_")
	fmt.Println(arr[0])
}

func excel() {
	// 打开XLSX文件
	f, err := excelize.OpenFile(path2)
	if err != nil {
		log.Fatal(err)
	}

	// 获取第一个表单
	sheetName := f.GetSheetName(1)
	if sheetName == "" {
		fmt.Println("can not get xlsx file sheetName")
		return
	}
	// 获取表单中的行数和列数
	rows := f.GetRows(sheetName)
	if len(rows) == 0 {
		fmt.Println("the file row is empty")
		return
	}
	header := make(map[int]string)
	for k, col := range rows[0] {
		if k >= 2 {
			header[k] = col
			fmt.Println(col)
		}
	}

	var parserMsg []Msg
	// 遍历行和列
	for k, row := range rows {
		if k == 0 {
			continue
		}

		var timeCol int64

		for ck, colCell := range row {
			if ck == 0 {
				parsedTime, err := time.Parse(laytime_local, colCell)
				if err != nil {
					fmt.Println("parse datapoint time failed:",err.Error())
					return
				}

				timeCol = parsedTime.Unix()
				continue
			}
			if ck >= 2 {
				value, err := strconv.ParseFloat(colCell, 64)
				if err != nil {
					fmt.Println("parser coll value failed:", err.Error())

					value = 0
				}
				parserMsg = append(parserMsg, Msg{
					Time:          timeCol,
					DatapointName: header[ck],
					Value:         value,
				})
			}
		}
	}


	for _, val := range parserMsg {
		fmt.Println("time:", val.Time, " DatapointName:", val.DatapointName, " Value:", val.Value)
	}
}