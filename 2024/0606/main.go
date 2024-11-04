package main

import (
	"fmt"
	"time"
)

const oneDaySeconds = 86400

func main() {
	startDate := int64(1714492800)
	endDate := int64(1717257600)
	//povID := "56575f1c-a9fb-4c29-8291-7c86742f73aa"
	//loc := "Asia/Hong_Kong"

	createAt := startDate

	bQuit := false
	for !bQuit {
		startTime := endDate - oneDaySeconds
		// rabbitmq.PublishEventEnergySummaryChanged(povID, startTime, endDate, timezoneRes.GetTimezone())
		fmt.Println("startTime:", startTime, "endTime:", endDate)
		fmt.Println("startDate:", time.Unix(startTime, 0).Format(time.DateTime), "endDate:", time.Unix(endDate, 0).Format(time.DateTime))
		if startTime < createAt {
			bQuit = true
		}

		endDate = startTime
	}
}
