package main

import (
	"fmt"
	"time"
)

func TimeSample() {
	now := time.Now()
	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)
	fmt.Println(now.String())
	//JavaScriptのDate.parse()でparse可能
	//JavaScriptのDate.toISOString()で生成した文字列をtime.RFC3339Nanoでパース可能
	fmt.Println(future.Format(time.RFC3339Nano))
}

func TimeDurationSample() {
	fiveMinute := time.Minute * 5
	var seconds int64 = 10
	tenSeconds := time.Duration(seconds) * time.Second
	fmt.Println(tenSeconds)
	past := time.Date(1955, time.November, 12, 6, 38, 0, 0, time.UTC)
	dur := time.Now().Sub(past)
	fmt.Println(dur)

	filePath := time.Now().Truncate(time.Hour).Format("20060102140405.json")
	fiveMinuteAfter := time.Now().Add(fiveMinute)
	fiveMinuteBefore := time.Now().Add(-fiveMinute)
	fmt.Println(filePath, fiveMinuteAfter.Unix, fiveMinuteBefore.Unix)
}

func NextMonth(t *time.Time) time.Time {
	year, month, day := t.Date()
	first := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	year2, month2, _ := first.AddDate(0, 1, 0).Date()
	nextMonthTime := time.Date(year2, month2, day, 0, 0, 0, 0, time.UTC)
	if nextMonthTime.Month() != month2 {
		return first.AddDate(0, 2, -1)
	}
	return nextMonthTime
}
