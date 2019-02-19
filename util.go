package main

import "time"

func convertDateToTime(zdate, ztime string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04", zdate+" "+ztime)
}
