package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间
	now := time.Now()

	// 今天开始时间和结束时间
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := todayStart.AddDate(0, 0, 1).Add(-time.Nanosecond)

	// 昨天开始时间和结束时间
	yesterdayStart := todayStart.AddDate(0, 0, -1)
	yesterdayEnd := yesterdayStart.AddDate(0, 0, 1).Add(-time.Nanosecond)

	// 7天开始时间和结束时间
	sevenDaysAgo := now.AddDate(0, 0, -7)
	sevenDaysStart := time.Date(sevenDaysAgo.Year(), sevenDaysAgo.Month(), sevenDaysAgo.Day(), 0, 0, 0, 0, sevenDaysAgo.Location())
	sevenDaysEnd := todayEnd

	// 一个月开始时间和结束时间
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneMonthStart := time.Date(oneMonthAgo.Year(), oneMonthAgo.Month(), oneMonthAgo.Day(), 0, 0, 0, 0, oneMonthAgo.Location())
	oneMonthEnd := todayEnd

	// 本年开始时间和结束时间
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	yearEnd := yearStart.AddDate(1, 0, 0).Add(-time.Nanosecond)

	// format
	format := "2006-01-02 15:04:05"
	fmt.Println("today:     ", todayStart.Format(format), todayEnd.Format(format))
	fmt.Println("yesterday: ", yesterdayStart.Format(format), yesterdayEnd.Format(format))
	fmt.Println("seven days:", sevenDaysStart.Format(format), sevenDaysEnd.Format(format))
	fmt.Println("one month: ", oneMonthStart.Format(format), oneMonthEnd.Format(format))
	fmt.Println("year:      ", yearStart.Format(format), yearEnd.Format(format))
}
