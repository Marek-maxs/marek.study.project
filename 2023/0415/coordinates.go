package main

import (
	"fmt"
	"math"
)

// 定义无人机坐标类型
type Coordinates struct {
	x float64
	y float64
	z float64
}

// 计算两个坐标之间的距离
func distance(coord1 Coordinates, coord2 Coordinates) float64 {
	dx := coord1.x - coord2.x
	dy := coord1.y - coord2.y
	dz := coord1.z - coord2.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func runCoordinates() {
	// 模拟无人机移动，获取当前坐标
	currentCoord := Coordinates{x: 10, y: 20, z: 5}
	fmt.Printf("当前无人机坐标：(%f, %f, %f)\n", currentCoord.x, currentCoord.y, currentCoord.z)

	// 定义目标坐标
	targetCoord := Coordinates{x: 100, y: 50, z: 10}
	fmt.Printf("目标无人机坐标：(%f, %f, %f)\n", targetCoord.x, targetCoord.y, targetCoord.z)

	// 计算无人机与目标之间的距离
	dist := distance(currentCoord, targetCoord)
	fmt.Printf("无人机与目标之间的距离为：%f\n", dist)
}

// 计算两点之间的距离
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const radius = 6371 // 地球半径，单位为千米
	dlat := (lat2 - lat1) * math.Pi / 180
	dlon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radius * c
}

// 计算定位结果
func getLocation(point1, point2, point3 [2]float64) (float64, float64) {
	// 计算三角形三边长
	a := distance(point1[0], point1[1], point2[0], point2[1])
	b := distance(point2[0], point2[1], point3[0], point3[1])
	c := distance(point3[0], point3[1], point1[0], point1[1])
	// 计算三角形内角余弦值
	cosa := (b*b + c*c - a*a) / (2 * b * c)
	cosb := (a*a + c*c - b*b) / (2 * a * c)
	cosc := (a*a + b*b - c*c) / (2 * a * b)
	// 计算三角形内角
	angA := math.Acos(cosa) * 180 / math.Pi
	angB := math.Acos(cosb) * 180 / math.Pi
	angC := math.Acos(cosc) * 180 / math.Pi
	// 根据三角形内角计算定位结果
	lat := (point1[0] + point2[0] + point3[0]) / 3
	lon := (point1[1] + point2[1] + point3[1]) / 3
	if angA > 90 {
		// test print value
		fmt.Println(angB, angC, lat, lon)
	}

	return 0, 0
}
