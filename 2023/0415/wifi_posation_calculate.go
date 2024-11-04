package main

import (
	"fmt"
	"math"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
基于WiFi信号强度的定位算法

在这里，我将演示一个基于WiFi信号强度的定位算法，它使用Go语言实现。该算法基于机器学习的原理，使用了一种称为K近邻（K-Nearest Neighbor）的算法。这种算法可以将位置信息转换为数字特征，然后找到离待定位点最近的已知位置点。

以下是算法的实现步骤：

1. 收集WiFi信号强度数据和位置信息，将它们存储在一个数据集中。
2. 对数据集进行预处理，例如将信号强度值进行标准化，使其具有

首先需要确定所需的定位算法类型，例如基于GPS、基于Wi-Fi信号强度或基于图像识别等。下面我将以基于Wi-Fi信号强度的定位算法为例，提供一个基本的Golang实现。
*/

/*
1. 收集Wi-Fi信号强度数据

首先需要通过程序收集Wi-Fi信号强度数据，可以使用Golang中的`os/exec`包执行系统命令来获取相关数据。以下是一个简单的代码示例：
*/

func wifiHightData() {
	cmd := exec.Command("iwlist", "wlan0", "scan")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile(`Address: ([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)
	macAddrs := re.FindAllStringSubmatch(string(out), -1)

	re = regexp.MustCompile(`Signal level=(-\d+) dBm`)
	signalLevels := re.FindAllStringSubmatch(string(out), -1)

	wifiSignals := make(map[string]int)
	for i, macAddr := range macAddrs {
		signalLevel, _ := strconv.Atoi(strings.Trim(signalLevels[i][1], " "))
		wifiSignals[macAddr[0][9:]] = signalLevel
	}

	fmt.Println(wifiSignals)
}

// 该代码解析了通过`iwlist wlan0 scan`命令获取到的Wi-Fi信号强度数据，将每个热点的信号强度值存储在了`wifiSignals`字典中。

/*
1. 定位算法实现

解析Wi-Fi信号强度数据后，可以使用定位算法来确定设备的位置。以下是一个简单的定位算法实现：
*/

type AccessPoint struct {
	X           float64
	Y           float64
	SignalLevel int
}

func runAccessPoint() {
	cmd := exec.Command("iwlist", "wlan0", "scan")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	// 解析Wi-Fi信号强度数据
	re := regexp.MustCompile(`Address: ([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)
	macAddrs := re.FindAllStringSubmatch(string(out), -1)

	re = regexp.MustCompile(`Signal level=(-\d+) dBm`)
	signalLevels := re.FindAllStringSubmatch(string(out), -1)

	wifiSignals := make(map[string]int)
	for i, macAddr := range macAddrs {
		signalLevel, _ := strconv.Atoi(strings.Trim(signalLevels[i][1], " "))
		wifiSignals[macAddr[0][9:]] = signalLevel
	}

	// 计算设备的位置
	accessPoints := map[string]AccessPoint{
		"00:11:22:33:44:55": AccessPoint{X: 0, Y: 0, SignalLevel: -40},
		"AA:BB:CC:DD:EE:FF": AccessPoint{X: 10, Y: 0, SignalLevel: -40},
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile(`Address: ([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)
	macAddrs := re.FindAllStringSubmatch(string(out), -1)

	re = regexp.MustCompile(`Signal level=(-\d+) dBm`)
	signalLevels := re.FindAllStringSubmatch(string(out), -1)

	wifiSignals := make(map[string]int)
	for i, macAddr := range macAddrs {
		signalLevel, _ := strconv.Atoi(strings.Trim(signalLevels[i][1], " "))
		wifiSignals[macAddr[0][9:]] = signalLevel
	}

	// 定位算法实现
	accessPoints := []AccessPoint{
		{0, 0, -40},
		{0, 10, -50},
		{10, 0, -60},
		{10, 10, -70},
	}
	estimatedX, estimatedY := trilateration(wifiSignals, accessPoints)
	fmt.Printf("Estimated position: (%.2f, %.2f)\n", estimatedX, estimatedY)

}

func trilateration(wifiSignals map[string]int, accessPoints []AccessPoint) (float64, float64) {
	// 按信号强度值从大到小排序
	sortedAPs := make([]AccessPoint, 0)
	for mac, signal := range wifiSignals {
		if ap, ok := findAP(mac, accessPoints); ok {
			ap.SignalLevel = signal
			sortedAPs = append(sortedAPs, ap)
		}
	}
	sortBySignalLevel(sortedAPs)
	// 使用三边定位算法计算估算位置
	xa, ya := sortedAPs[0].X, sortedAPs[0].Y
	xb, yb := sortedAPs[1].X, sortedAPs[1].Y
	xc, yc := sortedAPs[2].X, sortedAPs[2].Y
	ra := distance(sortedAPs[0].SignalLevel)
	rb := distance(sortedAPs[1].SignalLevel)
	rc := distance(sortedAPs[2].SignalLevel)
	P1 := xa*xa + ya*ya - ra*ra
	P2 := xb*xb + yb*yb - rb*rb
	P3 := xc*xc + yc*yc - rc*rc
	denominator := 2 * (xa*(yb-yc) + xb*(yc-ya) + xc*(ya-yb))
	estimatedX := (P1*(yb-yc) + P2*(yc-ya) + P3*(ya-yb)) / denominator
	estimatedY := (P1*(xc-xb) + P2*(xa-xc) + P3*(xb-xa)) / denominator
	return estimatedX, estimatedY
}

func findAP(mac string, accessPoints []AccessPoint) (AccessPoint, bool) {
	for _, ap := range accessPoints {
		if mac == fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X",
			byte(ap.X), byte(ap.X>>8), byte(ap.Y), byte(ap.Y>>8), 0x01, 0x01) {
			return ap, true
		}
	}
	return AccessPoint{}, false
}

func sortBySignalLevel(aps []AccessPoint) {
	// 信号强度值从大到小排序
	sort.SliceStable(aps, func(i, j int) bool {
		return aps[i].SignalLevel > aps[j].SignalLevel
	})
}

func distance(signalLevel int) float64 {
	return math.Pow(10, float64(-signalLevel-40)/20)
}

type AccessPoint struct {
	X           float64
	Y           float64
	SignalLevel int
}
