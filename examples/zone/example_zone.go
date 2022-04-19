package main

import (
	"fmt"
	"suzaku/pkg/utils"
	"time"
)

func main() {
	var (
		zone         int
		now          time.Time
		location     *time.Location
		timeStr      string
		locationTime time.Time

		utcStr string
		val    float64
	)

	now = time.Now()
	for i := -12; i <= 12; i++ {
		location = utils.FixedZone(i * 3600)
		timeStr = now.In(location).Format("2006-01-02 15:04:05")
		locationTime, _ = utils.ToLocationTime("2022-03-31", "2006-01-02", location)
		fmt.Println(locationTime, " | ", locationTime.In(utils.BeijingLocation()), " | ", "UTC", i)
		zone += 1
	}

	fmt.Println("---------------------------------")

	for i := -12; i <= 12; i++ {
		utcStr = fmt.Sprintf("UTC%d", i)
		location = utils.FixedZone(int(utils.UtcOffset(utcStr)))
		timeStr = now.In(location).Format("2006-01-02")
		fmt.Println(timeStr, utcStr, val)
	}
}
