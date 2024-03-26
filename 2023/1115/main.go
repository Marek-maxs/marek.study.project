package main

import (
	"fmt"
	"sort"
	"time"
)

/**
*
* Author: Marek
* Date: 2023-11-15 12:23
* Email: 364021318@qq.com
*
 */

type Device struct {
	DeviceID     string
	Gateway      string
	Serial       string
	Provider     string
	FanSpeed     string
	DeviceType   string
	Status       int
	DimmingLevel int
	HeatMode     int
	Temperature  float32
}

type sortByGateway []Device

func (s sortByGateway) Len() int           { return len(s) }
func (s sortByGateway) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByGateway) Less(i, j int) bool { return s[i].Gateway < s[j].Gateway }

func GatewayGrouping(list []Device) [][]Device {
	sort.Sort(sortByGateway(list))

	returnData := make([][]Device, 0)
	i := 0
	var j int

	for {
		if i >= len(list) {
			break
		}

		for j = i + 1; j < len(list) && list[i].Gateway == list[j].Gateway; j++ {
		}

		returnData = append(returnData, list[i:j])

		i = j
	}

	return returnData
}

func main() {
	dev := make([]Device, 0,5)

	dev = append(dev, Device{Gateway:"3C6A2CFFFED020A0", Serial:"E874D0FEFF2C6A3C1"})
	dev = append(dev, Device{Gateway:"3C6A2CFFFED020A0", Serial:"E874D0FEFF2C6A3C6"})
	dev = append(dev, Device{Gateway:"3C6A2CFFFED020A5", Serial:"AB74D0FEFF2C6A3C1"})
	dev = append(dev, Device{Gateway:"3C6A2CFFFED02111", Serial:"D574D0FEFF2C6A3C1"})
	dev = append(dev, Device{Gateway:"3C6A2CFFFED0210E", Serial:"F874D0FEFF2C6A3C1"})

	list := GatewayGrouping(dev)

	for k, v := range list{
		fmt.Println(k,":", v)
		for _, vv := range v {
			fmt.Println(vv.Serial)
		}
	}

	loc, _ := time.LoadLocation("Asia/Hong_Kong")
	startTime := 1700582400
	localTime := time.Unix(int64(startTime), 0).In(loc)
	fmt.Println(localTime.Format(time.RFC3339))
}