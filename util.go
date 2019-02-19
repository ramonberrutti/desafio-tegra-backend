package main

import "time"

func convertDateToTime(zdate, ztime string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04", zdate+" "+ztime)
}

func compareDate(time1, time2 time.Time) bool {
	return time1.Year() == time2.Year() && time1.Month() == time2.Month() && time1.Day() == time2.Day()
}
