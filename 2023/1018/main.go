package main

import (
	"fmt"
	"time"
)

const NewDateTime = "2006-01-02 15:04"

func main() {
fmt.Println(time.Now().Format(NewDateTime))

	if str := true; str {
		fmt.Println("sss")

}