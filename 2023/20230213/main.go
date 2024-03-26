package main

import (
	"fmt"
	"time"
)

var UnconfiguredDatapointMap = map[string]string{}

func main () {
for i := 0; i < 10 ; i++ {
	UnconfiguredDatapointMap[fmt.Sprintf("%d", i)] = "ss"
	fmt.Println(time.Now().Hour())
}
for _, val := range UnconfiguredDatapointMap{
	fmt.Println(val)
}
UnconfiguredDatapointMap = make(map[string]string)
fmt.Println(UnconfiguredDatapointMap)
}