package main

import "fmt"

var roundsArray = [...]string{
	0: "☆☆☆☆☆☆☆☆☆☆",
	1: "★☆☆☆☆☆☆☆☆☆",
	2: "★★☆☆☆☆☆☆☆☆",
	3: "★★★☆☆☆☆☆☆☆",
	4: "★★★★☆☆☆☆☆☆",
	5: "★★★★★☆☆☆☆☆",
	6: "★★★★★★☆☆☆☆",
	7: "★★★★★★★☆☆☆",
	8: "★★★★★★★★☆☆",
	9: "★★★★★★★★★☆",
}

func getPercentageRoundsArray(percentage float64) string {
	if p := int(percentage * 10); p >= 0 && p < len(roundsArray) {
		return roundsArray[p]
	} else {
		return "★★★★★★★★★★"
	}
}

var roundsMap = map[int]string{
	0: "☆☆☆☆☆☆☆☆☆☆",
	1: "★☆☆☆☆☆☆☆☆☆",
	2: "★★☆☆☆☆☆☆☆☆",
	3: "★★★☆☆☆☆☆☆☆",
	4: "★★★★☆☆☆☆☆☆",
	5: "★★★★★☆☆☆☆☆",
	6: "★★★★★★☆☆☆☆",
	7: "★★★★★★★☆☆☆",
	8: "★★★★★★★★☆☆",
	9: "★★★★★★★★★☆",
}

func getPercentageRoundsMap(percentage float64) string {
	if p, ok := roundsMap[int(percentage * 10)]; ok {
		return p
	} else {
		return "★★★★★★★★★★"
	}
}

func main() {
	var k float64
	for i := 0; i < 10; i++ {
		res := getPercentageRoundsArray(k)
		fmt.Println(res)
		k = k + 0.1
	}
	str := uui
}